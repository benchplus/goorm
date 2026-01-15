package sqlx

import (
	"os"

	"github.com/benchplus/goorm/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type SqlxORM struct {
	db *sqlx.DB
}

func New() *SqlxORM {
	return &SqlxORM{}
}

func (s *SqlxORM) Init(dsn string) error {
	var err error
	s.db, err = sqlx.Connect("sqlite3", dsn)
	return err
}

func (s *SqlxORM) Close() error {
	return s.db.Close()
}

func (s *SqlxORM) CreateTable() error {
	// 创建 users 表
	_, err := s.db.Exec(`
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
	_, err = s.db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			title VARCHAR(200) NOT NULL,
			body TEXT NOT NULL
		)
	`)
	return err
}

func (s *SqlxORM) DropTable() error {
	_, err := s.db.Exec("DROP TABLE IF EXISTS users")
	if err != nil {
		return err
	}
	_, err = s.db.Exec("DROP TABLE IF EXISTS posts")
	return err
}

func (s *SqlxORM) Insert(user *models.User) error {
	query := `INSERT INTO users (name, email, age) VALUES (?, ?, ?)`
	result, err := s.db.Exec(query, user.Name, user.Email, user.Age)
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

func (s *SqlxORM) InsertBatch(users []*models.User) error {
	query := `INSERT INTO users (name, email, age) VALUES (?, ?, ?)`
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Preparex(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, user := range users {
		result, err := stmt.Exec(user.Name, user.Email, user.Age)
		if err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		user.ID = id
	}
	return tx.Commit()
}

func (s *SqlxORM) GetByID(id int64) (*models.User, error) {
	user := &models.User{}
	err := s.db.Get(user, "SELECT id, name, email, age FROM users WHERE id = ?", id)
	return user, err
}

func (s *SqlxORM) GetByIDs(ids []int64) ([]*models.User, error) {
	query, args, err := sqlx.In("SELECT id, name, email, age FROM users WHERE id IN (?)", ids)
	if err != nil {
		return nil, err
	}
	query = s.db.Rebind(query)

	var users []*models.User
	err = s.db.Select(&users, query, args...)
	return users, err
}

func (s *SqlxORM) Update(user *models.User) error {
	query := `UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?`
	_, err := s.db.Exec(query, user.Name, user.Email, user.Age, user.ID)
	return err
}

func (s *SqlxORM) Delete(id int64) error {
	_, err := s.db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}

func (s *SqlxORM) Count() (int64, error) {
	var count int64
	err := s.db.Get(&count, "SELECT COUNT(*) FROM users")
	return count, err
}

func (s *SqlxORM) GetAll(limit, offset int) ([]*models.User, error) {
	var users []*models.User
	err := s.db.Select(&users, "SELECT id, name, email, age FROM users LIMIT ? OFFSET ?", limit, offset)
	return users, err
}

// GetDSN 生成测试用的 DSN
func GetDSN() string {
	return getTempFile()
}

func getTempFile() string {
	tmpfile, _ := os.CreateTemp("", "sqlx_*.db")
	tmpfile.Close()
	return tmpfile.Name()
}
