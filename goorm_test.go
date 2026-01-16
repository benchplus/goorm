package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/benchplus/goorm/borm"
	"github.com/benchplus/goorm/bun"
	"github.com/benchplus/goorm/ent"
	"github.com/benchplus/goorm/gorm"
	"github.com/benchplus/goorm/internal/models"
	"github.com/benchplus/goorm/internal/orm"
	"github.com/benchplus/goorm/sqlx"
	"github.com/benchplus/goorm/xorm"
	"github.com/benchplus/goorm/zorm"
)

// 测试用的 ORM 实现列表
var orms = map[string]struct {
	init func() orm.Interface
	dsn  func() string
}{
	"gorm": {
		init: func() orm.Interface { return gorm.New() },
		dsn:  gorm.GetDSN,
	},
	"xorm": {
		init: func() orm.Interface { return xorm.New() },
		dsn:  xorm.GetDSN,
	},
	"zorm": {
		init: func() orm.Interface { return zorm.New() },
		dsn:  zorm.GetDSN,
	},
	"sqlx": {
		init: func() orm.Interface { return sqlx.New() },
		dsn:  sqlx.GetDSN,
	},
	"borm": {
		init: func() orm.Interface { return borm.New() },
		dsn:  borm.GetDSN,
	},
	"bun": {
		init: func() orm.Interface { return bun.New() },
		dsn:  bun.GetDSN,
	},
	"ent": {
		init: func() orm.Interface { return ent.New() },
		dsn:  ent.GetDSN,
	},
}

// setupORM 初始化 ORM
func setupORM(name string) (orm.Interface, func(), error) {
	ormInfo, ok := orms[name]
	if !ok {
		return nil, nil, fmt.Errorf("unknown ORM: %s", name)
	}

	orm := ormInfo.init()
	dsn := ormInfo.dsn()

	if err := orm.Init(dsn); err != nil {
		return nil, nil, err
	}

	if err := orm.CreateTable(); err != nil {
		orm.Close()
		return nil, nil, err
	}

	cleanup := func() {
		orm.DropTable()
		orm.Close()
		// 清理临时文件
		if dsn != "" {
			os.Remove(dsn)
		}
	}

	return orm, cleanup, nil
}

// BenchmarkInsertSingle 单条插入测试
func BenchmarkInsertSingle_GORM(b *testing.B) {
	benchmarkInsertSingle(b, "gorm")
}

func BenchmarkInsertSingle_XORM(b *testing.B) {
	benchmarkInsertSingle(b, "xorm")
}

func BenchmarkInsertSingle_ZORM(b *testing.B) {
	benchmarkInsertSingle(b, "zorm")
}

func BenchmarkInsertSingle_SQLX(b *testing.B) {
	benchmarkInsertSingle(b, "sqlx")
}

func BenchmarkInsertSingle_BORM(b *testing.B) {
	benchmarkInsertSingle(b, "borm")
}

func BenchmarkInsertSingle_BUN(b *testing.B) {
	benchmarkInsertSingle(b, "bun")
}

func BenchmarkInsertSingle_ENT(b *testing.B) {
	benchmarkInsertSingle(b, "ent")
}

func benchmarkInsertSingle(b *testing.B, ormName string) {
	orm, cleanup, err := setupORM(ormName)
	if err != nil {
		b.Fatalf("Setup failed: %v", err)
	}
	defer cleanup()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		user := &models.User{
			Name:  fmt.Sprintf("user%d", i),
			Email: fmt.Sprintf("user%d@example.com", i),
			Age:   20 + (i % 50),
		}
		if err := orm.Insert(user); err != nil {
			b.Fatalf("Insert failed: %v", err)
		}
	}
}

// BenchmarkInsertBatch 批量插入测试
func BenchmarkInsertBatch_GORM(b *testing.B) {
	benchmarkInsertBatch(b, "gorm")
}

func BenchmarkInsertBatch_XORM(b *testing.B) {
	benchmarkInsertBatch(b, "xorm")
}

func BenchmarkInsertBatch_ZORM(b *testing.B) {
	benchmarkInsertBatch(b, "zorm")
}

func BenchmarkInsertBatch_SQLX(b *testing.B) {
	benchmarkInsertBatch(b, "sqlx")
}

func BenchmarkInsertBatch_BORM(b *testing.B) {
	benchmarkInsertBatch(b, "borm")
}

func BenchmarkInsertBatch_BUN(b *testing.B) {
	benchmarkInsertBatch(b, "bun")
}

func BenchmarkInsertBatch_ENT(b *testing.B) {
	benchmarkInsertBatch(b, "ent")
}

func benchmarkInsertBatch(b *testing.B, ormName string) {
	orm, cleanup, err := setupORM(ormName)
	if err != nil {
		b.Fatalf("Setup failed: %v", err)
	}
	defer cleanup()

	batchSize := 100
	users := make([]*models.User, batchSize)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for j := 0; j < batchSize; j++ {
			users[j] = &models.User{
				Name:  fmt.Sprintf("user%d_%d", i, j),
				Email: fmt.Sprintf("user%d_%d@example.com", i, j),
				Age:   20 + (j % 50),
			}
		}
		if err := orm.InsertBatch(users); err != nil {
			b.Fatalf("InsertBatch failed: %v", err)
		}
	}
}

// BenchmarkGetByID 根据 ID 查询测试
func BenchmarkGetByID_GORM(b *testing.B) {
	benchmarkGetByID(b, "gorm")
}

func BenchmarkGetByID_XORM(b *testing.B) {
	benchmarkGetByID(b, "xorm")
}

func BenchmarkGetByID_ZORM(b *testing.B) {
	benchmarkGetByID(b, "zorm")
}

func BenchmarkGetByID_SQLX(b *testing.B) {
	benchmarkGetByID(b, "sqlx")
}

func BenchmarkGetByID_BORM(b *testing.B) {
	benchmarkGetByID(b, "borm")
}

func BenchmarkGetByID_BUN(b *testing.B) {
	benchmarkGetByID(b, "bun")
}

func BenchmarkGetByID_ENT(b *testing.B) {
	benchmarkGetByID(b, "ent")
}

func benchmarkGetByID(b *testing.B, ormName string) {
	orm, cleanup, err := setupORM(ormName)
	if err != nil {
		b.Fatalf("Setup failed: %v", err)
	}
	defer cleanup()

	// 预先插入一些数据
	var ids []int64
	for i := 0; i < 1000; i++ {
		user := &models.User{
			Name:  fmt.Sprintf("user%d", i),
			Email: fmt.Sprintf("user%d@example.com", i),
			Age:   20 + (i % 50),
		}
		if err := orm.Insert(user); err != nil {
			b.Fatalf("Pre-insert failed: %v", err)
		}
		ids = append(ids, user.ID)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		id := ids[i%len(ids)]
		_, err := orm.GetByID(id)
		if err != nil {
			b.Fatalf("GetByID failed: %v", err)
		}
	}
}

// BenchmarkGetByIDs 根据多个 ID 查询测试
func BenchmarkGetByIDs_GORM(b *testing.B) {
	benchmarkGetByIDs(b, "gorm")
}

func BenchmarkGetByIDs_XORM(b *testing.B) {
	benchmarkGetByIDs(b, "xorm")
}

func BenchmarkGetByIDs_ZORM(b *testing.B) {
	benchmarkGetByIDs(b, "zorm")
}

func BenchmarkGetByIDs_SQLX(b *testing.B) {
	benchmarkGetByIDs(b, "sqlx")
}

func BenchmarkGetByIDs_BORM(b *testing.B) {
	benchmarkGetByIDs(b, "borm")
}

func BenchmarkGetByIDs_BUN(b *testing.B) {
	benchmarkGetByIDs(b, "bun")
}

func BenchmarkGetByIDs_ENT(b *testing.B) {
	benchmarkGetByIDs(b, "ent")
}

func benchmarkGetByIDs(b *testing.B, ormName string) {
	orm, cleanup, err := setupORM(ormName)
	if err != nil {
		b.Fatalf("Setup failed: %v", err)
	}
	defer cleanup()

	// 预先插入一些数据
	var ids []int64
	for i := 0; i < 1000; i++ {
		user := &models.User{
			Name:  fmt.Sprintf("user%d", i),
			Email: fmt.Sprintf("user%d@example.com", i),
			Age:   20 + (i % 50),
		}
		if err := orm.Insert(user); err != nil {
			b.Fatalf("Pre-insert failed: %v", err)
		}
		ids = append(ids, user.ID)
	}

	batchSize := 10
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		start := (i * batchSize) % len(ids)
		end := start + batchSize
		if end > len(ids) {
			end = len(ids)
		}
		batchIDs := ids[start:end]
		_, err := orm.GetByIDs(batchIDs)
		if err != nil {
			b.Fatalf("GetByIDs failed: %v", err)
		}
	}
}

// BenchmarkUpdate 更新测试
func BenchmarkUpdate_GORM(b *testing.B) {
	benchmarkUpdate(b, "gorm")
}

func BenchmarkUpdate_XORM(b *testing.B) {
	benchmarkUpdate(b, "xorm")
}

func BenchmarkUpdate_ZORM(b *testing.B) {
	benchmarkUpdate(b, "zorm")
}

func BenchmarkUpdate_SQLX(b *testing.B) {
	benchmarkUpdate(b, "sqlx")
}

func BenchmarkUpdate_BORM(b *testing.B) {
	benchmarkUpdate(b, "borm")
}

func BenchmarkUpdate_BUN(b *testing.B) {
	benchmarkUpdate(b, "bun")
}

func BenchmarkUpdate_ENT(b *testing.B) {
	benchmarkUpdate(b, "ent")
}

func benchmarkUpdate(b *testing.B, ormName string) {
	orm, cleanup, err := setupORM(ormName)
	if err != nil {
		b.Fatalf("Setup failed: %v", err)
	}
	defer cleanup()

	// 预先插入一些数据
	var users []*models.User
	for i := 0; i < 1000; i++ {
		user := &models.User{
			Name:  fmt.Sprintf("user%d", i),
			Email: fmt.Sprintf("user%d@example.com", i),
			Age:   20 + (i % 50),
		}
		if err := orm.Insert(user); err != nil {
			b.Fatalf("Pre-insert failed: %v", err)
		}
		users = append(users, user)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		user := users[i%len(users)]
		user.Name = fmt.Sprintf("updated_user%d", i)
		user.Age = 30 + (i % 50)
		if err := orm.Update(user); err != nil {
			b.Fatalf("Update failed: %v", err)
		}
	}
}

// BenchmarkDelete 删除测试
func BenchmarkDelete_GORM(b *testing.B) {
	benchmarkDelete(b, "gorm")
}

func BenchmarkDelete_XORM(b *testing.B) {
	benchmarkDelete(b, "xorm")
}

func BenchmarkDelete_ZORM(b *testing.B) {
	benchmarkDelete(b, "zorm")
}

func BenchmarkDelete_SQLX(b *testing.B) {
	benchmarkDelete(b, "sqlx")
}

func BenchmarkDelete_BORM(b *testing.B) {
	benchmarkDelete(b, "borm")
}

func BenchmarkDelete_BUN(b *testing.B) {
	benchmarkDelete(b, "bun")
}

func BenchmarkDelete_ENT(b *testing.B) {
	benchmarkDelete(b, "ent")
}

func benchmarkDelete(b *testing.B, ormName string) {
	orm, cleanup, err := setupORM(ormName)
	if err != nil {
		b.Fatalf("Setup failed: %v", err)
	}
	defer cleanup()

	// 预先插入大量数据
	var ids []int64
	for i := 0; i < b.N+1000; i++ {
		user := &models.User{
			Name:  fmt.Sprintf("user%d", i),
			Email: fmt.Sprintf("user%d@example.com", i),
			Age:   20 + (i % 50),
		}
		if err := orm.Insert(user); err != nil {
			b.Fatalf("Pre-insert failed: %v", err)
		}
		ids = append(ids, user.ID)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := orm.Delete(ids[i]); err != nil {
			b.Fatalf("Delete failed: %v", err)
		}
	}
}

// BenchmarkCount 统计数量测试
func BenchmarkCount_GORM(b *testing.B) {
	benchmarkCount(b, "gorm")
}

func BenchmarkCount_XORM(b *testing.B) {
	benchmarkCount(b, "xorm")
}

func BenchmarkCount_ZORM(b *testing.B) {
	benchmarkCount(b, "zorm")
}

func BenchmarkCount_SQLX(b *testing.B) {
	benchmarkCount(b, "sqlx")
}

func BenchmarkCount_BORM(b *testing.B) {
	benchmarkCount(b, "borm")
}

func BenchmarkCount_BUN(b *testing.B) {
	benchmarkCount(b, "bun")
}

func BenchmarkCount_ENT(b *testing.B) {
	benchmarkCount(b, "ent")
}

func benchmarkCount(b *testing.B, ormName string) {
	orm, cleanup, err := setupORM(ormName)
	if err != nil {
		b.Fatalf("Setup failed: %v", err)
	}
	defer cleanup()

	// 预先插入一些数据
	for i := 0; i < 1000; i++ {
		user := &models.User{
			Name:  fmt.Sprintf("user%d", i),
			Email: fmt.Sprintf("user%d@example.com", i),
			Age:   20 + (i % 50),
		}
		if err := orm.Insert(user); err != nil {
			b.Fatalf("Pre-insert failed: %v", err)
		}
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := orm.Count()
		if err != nil {
			b.Fatalf("Count failed: %v", err)
		}
	}
}

// BenchmarkGetAll 获取所有记录测试
func BenchmarkGetAll_GORM(b *testing.B) {
	benchmarkGetAll(b, "gorm")
}

func BenchmarkGetAll_XORM(b *testing.B) {
	benchmarkGetAll(b, "xorm")
}

func BenchmarkGetAll_ZORM(b *testing.B) {
	benchmarkGetAll(b, "zorm")
}

func BenchmarkGetAll_SQLX(b *testing.B) {
	benchmarkGetAll(b, "sqlx")
}

func BenchmarkGetAll_BORM(b *testing.B) {
	benchmarkGetAll(b, "borm")
}

func BenchmarkGetAll_BUN(b *testing.B) {
	benchmarkGetAll(b, "bun")
}

func BenchmarkGetAll_ENT(b *testing.B) {
	benchmarkGetAll(b, "ent")
}

func benchmarkGetAll(b *testing.B, ormName string) {
	orm, cleanup, err := setupORM(ormName)
	if err != nil {
		b.Fatalf("Setup failed: %v", err)
	}
	defer cleanup()

	// 预先插入一些数据
	for i := 0; i < 1000; i++ {
		user := &models.User{
			Name:  fmt.Sprintf("user%d", i),
			Email: fmt.Sprintf("user%d@example.com", i),
			Age:   20 + (i % 50),
		}
		if err := orm.Insert(user); err != nil {
			b.Fatalf("Pre-insert failed: %v", err)
		}
	}

	limit := 100
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		offset := (i * limit) % 900
		_, err := orm.GetAll(limit, offset)
		if err != nil {
			b.Fatalf("GetAll failed: %v", err)
		}
	}
}
