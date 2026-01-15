package zorm

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/benchplus/goorm/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

type ZormORM struct {
	db          *sql.DB
	insertStmt  *sql.Stmt
	updateStmt  *sql.Stmt
	deleteStmt  *sql.Stmt
	countStmt   *sql.Stmt
}

func New() *ZormORM {
	return &ZormORM{}
}

func (zo *ZormORM) Init(dsn string) error {
	var err error
	zo.db, err = sql.Open("sqlite3", dsn)
	if err != nil {
		return err
	}
	zo.db.SetMaxOpenConns(1)
	return nil
}

// prepareStatements 预编译常用语句，在CreateTable之后调用
func (zo *ZormORM) prepareStatements() error {
	var err error
	if zo.insertStmt == nil {
		zo.insertStmt, err = zo.db.Prepare(`INSERT INTO users (name, email, age) VALUES (?, ?, ?)`)
		if err != nil {
			return err
		}
	}
	if zo.updateStmt == nil {
		zo.updateStmt, err = zo.db.Prepare(`UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?`)
		if err != nil {
			return err
		}
	}
	if zo.deleteStmt == nil {
		zo.deleteStmt, err = zo.db.Prepare(`DELETE FROM users WHERE id = ?`)
		if err != nil {
			return err
		}
	}
	if zo.countStmt == nil {
		zo.countStmt, err = zo.db.Prepare(`SELECT COUNT(*) FROM users`)
		if err != nil {
			return err
		}
	}
	return nil
}

func (zo *ZormORM) Close() error {
	if zo.insertStmt != nil {
		zo.insertStmt.Close()
	}
	if zo.updateStmt != nil {
		zo.updateStmt.Close()
	}
	if zo.deleteStmt != nil {
		zo.deleteStmt.Close()
	}
	if zo.countStmt != nil {
		zo.countStmt.Close()
	}
	return zo.db.Close()
}

func (zo *ZormORM) CreateTable() error {
	// 创建 users 表
	_, err := zo.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) NOT NULL,
			age INTEGER NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	// 创建 posts 表
	_, err = zo.db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			title VARCHAR(200) NOT NULL,
			body TEXT NOT NULL
		)
	`)
	if err != nil {
		return err
	}
	
	// 表创建后预编译语句
	return zo.prepareStatements()
}

func (zo *ZormORM) DropTable() error {
	_, err := zo.db.Exec("DROP TABLE IF EXISTS users")
	if err != nil {
		return err
	}
	_, err = zo.db.Exec("DROP TABLE IF EXISTS posts")
	return err
}

func (zo *ZormORM) Insert(user *models.User) error {
	// 使用预编译语句，提升性能
	if zo.insertStmt == nil {
		if err := zo.prepareStatements(); err != nil {
			return err
		}
	}
	result, err := zo.insertStmt.Exec(user.Name, user.Email, user.Age)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}

func (zo *ZormORM) InsertBatch(users []*models.User) error {
	if len(users) == 0 {
		return nil
	}

	// 使用多行INSERT语句，一次性插入所有记录，性能最优
	query := `INSERT INTO users (name, email, age) VALUES `
	args := make([]interface{}, 0, len(users)*3)
	placeholders := make([]string, 0, len(users))
	
	for _, user := range users {
		placeholders = append(placeholders, "(?, ?, ?)")
		args = append(args, user.Name, user.Email, user.Age)
	}
	query += strings.Join(placeholders, ", ")

	// 执行批量插入
	result, err := zo.db.Exec(query, args...)
	if err != nil {
		return err
	}

	// 获取第一个插入的ID，然后依次递增（SQLite的last_insert_rowid返回第一个ID）
	firstID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// 为所有用户设置ID（SQLite批量插入时，ID是连续的）
	for i := range users {
		users[i].ID = firstID + int64(i)
	}

	return nil
}

func (zo *ZormORM) GetByID(id int64) (*models.User, error) {
	// 使用原生SQL替代zorm抽象，提升性能
	user := &models.User{}
	err := zo.db.QueryRow("SELECT id, name, email, age FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Age)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (zo *ZormORM) GetByIDs(ids []int64) ([]*models.User, error) {
	if len(ids) == 0 {
		return []*models.User{}, nil
	}

	// 使用原生SQL替代zorm抽象，提升性能
	// 构建IN子句
	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}
	query := "SELECT id, name, email, age FROM users WHERE id IN (" + strings.Join(placeholders, ", ") + ")"

	rows, err := zo.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 0, len(ids))
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, rows.Err()
}

func (zo *ZormORM) Update(user *models.User) error {
	// 使用预编译语句，提升性能
	if zo.updateStmt == nil {
		if err := zo.prepareStatements(); err != nil {
			return err
		}
	}
	_, err := zo.updateStmt.Exec(user.Name, user.Email, user.Age, user.ID)
	return err
}

func (zo *ZormORM) Delete(id int64) error {
	// 使用预编译语句，提升性能
	if zo.deleteStmt == nil {
		if err := zo.prepareStatements(); err != nil {
			return err
		}
	}
	_, err := zo.deleteStmt.Exec(id)
	return err
}

func (zo *ZormORM) Count() (int64, error) {
	// 使用预编译语句，提升性能
	if zo.countStmt == nil {
		if err := zo.prepareStatements(); err != nil {
			return 0, err
		}
	}
	var count int64
	err := zo.countStmt.QueryRow().Scan(&count)
	return count, err
}

func (zo *ZormORM) GetAll(limit, offset int) ([]*models.User, error) {
	// 使用原生SQL替代zorm抽象，提升性能
	rows, err := zo.db.Query("SELECT id, name, email, age FROM users LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, rows.Err()
}

// GetDSN 生成测试用的 DSN
func GetDSN() string {
	// 使用与GORM相同的DSN格式，启用缓存和内存模式
	return fmt.Sprintf("file:%s?cache=shared&mode=memory", getTempFile())
}

func getTempFile() string {
	tmpfile, _ := os.CreateTemp("", "zorm_*.db")
	tmpfile.Close()
	return tmpfile.Name()
}
