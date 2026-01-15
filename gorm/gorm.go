package gorm

import (
	"fmt"
	"os"

	"github.com/benchplus/goorm/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GormORM struct {
	db *gorm.DB
}

func New() *GormORM {
	return &GormORM{}
}

func (g *GormORM) Init(dsn string) error {
	var err error
	g.db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	return err
}

func (g *GormORM) Close() error {
	sqlDB, err := g.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (g *GormORM) CreateTable() error {
	return g.db.AutoMigrate(&models.User{})
}

func (g *GormORM) DropTable() error {
	return g.db.Migrator().DropTable(&models.User{})
}

func (g *GormORM) Insert(user *models.User) error {
	return g.db.Create(user).Error
}

func (g *GormORM) InsertBatch(users []*models.User) error {
	return g.db.CreateInBatches(users, 100).Error
}

func (g *GormORM) GetByID(id int64) (*models.User, error) {
	var user models.User
	err := g.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (g *GormORM) GetByIDs(ids []int64) ([]*models.User, error) {
	var users []*models.User
	err := g.db.Where("id IN ?", ids).Find(&users).Error
	return users, err
}

func (g *GormORM) Update(user *models.User) error {
	return g.db.Save(user).Error
}

func (g *GormORM) Delete(id int64) error {
	return g.db.Delete(&models.User{}, id).Error
}

func (g *GormORM) Count() (int64, error) {
	var count int64
	err := g.db.Model(&models.User{}).Count(&count).Error
	return count, err
}

func (g *GormORM) GetAll(limit, offset int) ([]*models.User, error) {
	var users []*models.User
	err := g.db.Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}

// GetDSN 生成测试用的 DSN
func GetDSN() string {
	return fmt.Sprintf("file:%s?cache=shared&mode=memory", getTempFile())
}

func getTempFile() string {
	tmpfile, _ := os.CreateTemp("", "gorm_*.db")
	tmpfile.Close()
	return tmpfile.Name()
}
