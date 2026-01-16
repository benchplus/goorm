package bun

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/benchplus/goorm/internal/models"
	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

type BunORM struct {
	db  *bun.DB
	ctx context.Context
}

func New() *BunORM {
	return &BunORM{
		ctx: context.Background(),
	}
}

func (b *BunORM) Init(dsn string) error {
	sqldb, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return err
	}
	sqldb.SetMaxOpenConns(1)
	b.db = bun.NewDB(sqldb, sqlitedialect.New())
	return nil
}

func (b *BunORM) Close() error {
	return b.db.Close()
}

func (b *BunORM) CreateTable() error {
	_, err := b.db.NewCreateTable().
		Model((*models.User)(nil)).
		IfNotExists().
		Exec(b.ctx)
	return err
}

func (b *BunORM) DropTable() error {
	_, err := b.db.NewDropTable().
		Model((*models.User)(nil)).
		IfExists().
		Exec(b.ctx)
	return err
}

func (b *BunORM) Insert(user *models.User) error {
	_, err := b.db.NewInsert().Model(user).Exec(b.ctx)
	if err != nil {
		return err
	}
	// BUN automatically fills the ID after insert if the model has pk tag
	return nil
}

func (b *BunORM) InsertBatch(users []*models.User) error {
	if len(users) == 0 {
		return nil
	}
	// 使用 Returning 来获取插入的 ID，避免 LastInsertId 不支持的问题
	err := b.db.NewInsert().
		Model(&users).
		Returning("id").
		Scan(b.ctx)
	if err != nil {
		return err
	}
	// BUN 会自动填充 users 中的 ID 字段
	return nil
}

func (b *BunORM) GetByID(id int64) (*models.User, error) {
	user := &models.User{}
	err := b.db.NewSelect().
		Model(user).
		Where("id = ?", id).
		Scan(b.ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (b *BunORM) GetByIDs(ids []int64) ([]*models.User, error) {
	var users []*models.User
	err := b.db.NewSelect().
		Model(&users).
		Where("id IN (?)", bun.In(ids)).
		Scan(b.ctx)
	return users, err
}

func (b *BunORM) Update(user *models.User) error {
	_, err := b.db.NewUpdate().
		Model(user).
		Where("id = ?", user.ID).
		Exec(b.ctx)
	return err
}

func (b *BunORM) Delete(id int64) error {
	_, err := b.db.NewDelete().
		Model((*models.User)(nil)).
		Where("id = ?", id).
		Exec(b.ctx)
	return err
}

func (b *BunORM) Count() (int64, error) {
	count, err := b.db.NewSelect().
		Model((*models.User)(nil)).
		Count(b.ctx)
	return int64(count), err
}

func (b *BunORM) GetAll(limit, offset int) ([]*models.User, error) {
	var users []*models.User
	err := b.db.NewSelect().
		Model(&users).
		Limit(limit).
		Offset(offset).
		Scan(b.ctx)
	return users, err
}

// GetDSN 生成测试用的 DSN
func GetDSN() string {
	return fmt.Sprintf("file:%s?cache=shared&mode=memory", getTempFile())
}

func getTempFile() string {
	tmpfile, _ := os.CreateTemp("", "bun_*.db")
	tmpfile.Close()
	return tmpfile.Name()
}
