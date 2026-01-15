package orm

import "github.com/benchplus/goorm/internal/models"

// Interface 统一的 ORM 接口
type Interface interface {
	// Init 初始化数据库连接
	Init(dsn string) error

	// Close 关闭数据库连接
	Close() error

	// CreateTable 创建表
	CreateTable() error

	// DropTable 删除表
	DropTable() error

	// Insert 插入单条记录
	Insert(user *models.User) error

	// InsertBatch 批量插入
	InsertBatch(users []*models.User) error

	// GetByID 根据 ID 查询
	GetByID(id int64) (*models.User, error)

	// GetByIDs 根据多个 ID 查询
	GetByIDs(ids []int64) ([]*models.User, error)

	// Update 更新记录
	Update(user *models.User) error

	// Delete 删除记录
	Delete(id int64) error

	// Count 统计数量
	Count() (int64, error)

	// GetAll 获取所有记录
	GetAll(limit, offset int) ([]*models.User, error)
}
