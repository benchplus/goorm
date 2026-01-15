package xorm

import (
	"fmt"
	"os"

	"github.com/benchplus/goorm/internal/models"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

type XormORM struct {
	engine *xorm.Engine
}

func New() *XormORM {
	return &XormORM{}
}

func (x *XormORM) Init(dsn string) error {
	var err error
	x.engine, err = xorm.NewEngine("sqlite3", dsn)
	if err != nil {
		return err
	}
	x.engine.SetMaxOpenConns(1)
	return nil
}

func (x *XormORM) Close() error {
	return x.engine.Close()
}

func (x *XormORM) CreateTable() error {
	return x.engine.Sync2(&models.User{})
}

func (x *XormORM) DropTable() error {
	return x.engine.DropTables(&models.User{})
}

func (x *XormORM) Insert(user *models.User) error {
	_, err := x.engine.Insert(user)
	return err
}

func (x *XormORM) InsertBatch(users []*models.User) error {
	_, err := x.engine.Insert(users)
	return err
}

func (x *XormORM) GetByID(id int64) (*models.User, error) {
	user := &models.User{}
	has, err := x.engine.ID(id).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (x *XormORM) GetByIDs(ids []int64) ([]*models.User, error) {
	var users []*models.User
	err := x.engine.In("id", ids).Find(&users)
	return users, err
}

func (x *XormORM) Update(user *models.User) error {
	_, err := x.engine.ID(user.ID).Update(user)
	return err
}

func (x *XormORM) Delete(id int64) error {
	_, err := x.engine.ID(id).Delete(&models.User{})
	return err
}

func (x *XormORM) Count() (int64, error) {
	return x.engine.Count(&models.User{})
}

func (x *XormORM) GetAll(limit, offset int) ([]*models.User, error) {
	var users []*models.User
	err := x.engine.Limit(limit, offset).Find(&users)
	return users, err
}

// GetDSN 生成测试用的 DSN
func GetDSN() string {
	return getTempFile()
}

func getTempFile() string {
	tmpfile, _ := os.CreateTemp("", "xorm_*.db")
	tmpfile.Close()
	return tmpfile.Name()
}
