# goorm

Go 语言 ORM 库持续性能基准测试

[English](README.md) | [中文](README_CN.md)

## 概述

本仓库使用 SQLite 作为测试数据库，为流行的 Go ORM 库提供性能基准测试。基准测试旨在比较不同 ORM 库在常见数据库操作上的性能。

## 测试的 ORM 库

- **[ZORM](https://github.com/IceWhaleTech/zorm)** - 简单、超快的 ORM 库
- **[BORM](https://github.com/orca-zhang/borm)** - 更好的 ORM 库，简单、快速且可自模拟
- **[GORM](https://gorm.io/)** - Go 最流行的 ORM 库
- **[SQLX](https://github.com/jmoiron/sqlx)** - 扩展 Go 标准 `database/sql` 库的库
- **[XORM](https://xorm.io/)** - 简单而强大的 Go ORM

## 基准测试

以下操作进行了基准测试：

| 测试用例 | 描述 |
|---------|------|
| `InsertSingle` | 单条记录插入性能 |
| `InsertBatch` | 批量插入性能（每批 100 条记录） |
| `GetByID` | 根据主键查询单条记录 |
| `GetByIDs` | 根据多个主键查询多条记录 |
| `Update` | 记录更新性能 |
| `Delete` | 记录删除性能 |
| `Count` | 统计查询性能 |
| `GetAll` | 分页查询性能（limit/offset） |

## 运行基准测试

### 前置要求

- Go 1.21 或更高版本
- SQLite3

### 运行所有基准测试

```bash
go test -bench=. -benchmem
```

### 运行特定 ORM 的基准测试

```bash
# 仅 GORM
go test -bench=GORM -benchmem

# 仅 XORM
go test -bench=XORM -benchmem

# 仅 ZORM
go test -bench=ZORM -benchmem

# 仅 SQLX
go test -bench=SQLX -benchmem

# 仅 BORM
go test -bench=BORM -benchmem
```

### 运行特定测试用例

```bash
# 所有 ORM 的 InsertSingle 测试
go test -bench=InsertSingle -benchmem

# 所有 ORM 的 GetByID 测试
go test -bench=GetByID -benchmem
```

## 基准测试结果

基准测试结果显示：
- **ns/op**: 每次操作的纳秒数
- **B/op**: 每次操作分配的字节数
- **allocs/op**: 每次操作的分配次数

### 快速摘要

| 测试用例 | [**ZORM**](https://github.com/IceWhaleTech/zorm) | [**BORM**](https://github.com/orca-zhang/borm) | [**GORM**](https://gorm.io/) | [**SQLX**](https://github.com/jmoiron/sqlx) | [**XORM**](https://xorm.io/) |
|---------|---------------------------------------------------|-------------------------------------------------|------------------------------|---------------------------------------------|------------------------------|
| InsertSingle | 🚀 | 🚀 | ✈️ | 🐌 | 🐌 |
| InsertBatch | 🚀 | 🚀 | ✈️ | 🐌 | 🐌 |
| GetByID | 🚀 | 🚀 | ✈️ | ✈️ | 🐌 |
| GetByIDs | 🚀 | 🚀 | ✈️ | ✈️ | 🐌 |
| Update | 🚀 | 🚀 | ✈️ | 🐌 | 🐌 |
| Delete | 🚀 | 🚀 | ✈️ | 🐌 | 🐌 |
| Count | 🚀 | 🚀 | ✈️ | 🐌 | 🐌 |
| GetAll | 🚀 | 🚀 | ✈️ | ✈️ | 🐌 |

> 🐌 表示非常慢, ✈️ 表示快, 🚀 表示非常快
>
> ⭐ 表示该 ORM 在该测试用例下 **同时更快且更省内存**（在 **ns/op** 与 **B/op** 这两个维度上为 Pareto 最优，数值越小越好）。⭐ 标记会打在 **ns/op** 和 **B/op** 两列中。

### 详细结果

#### 测试环境

- **Go Version**: 1.21+
- **Database**: SQLite（内存数据库，DSN 使用 `cache=shared&mode=memory`）
- **CPU**: M4 Pro
- **OS**: darwin (amd64)

#### InsertSingle（单条记录插入）

| ORM | ns/op | B/op | allocs/op |
|-----|-------|------|-----------|
| BORM | 2,970 ⭐ | 512 ⭐ | 22 |
| ZORM | 3,439 ⭐ | 512 ⭐ | 22 |
| GORM | 19,041 | 6,384 | 125 |
| SQLX | 183,124 | 692 ⭐ | 26 |
| XORM | 193,387 | 2,749 | 62 |

#### InsertBatch（批量插入 - 每批 100 条）

| ORM | ns/op | B/op | allocs/op |
|-----|-------|------|-----------|
| BORM | 101,798 ⭐ | 61,914 ⭐ | 1,216 |
| ZORM | 103,583 ⭐ | 61,913 ⭐ | 1,216 |
| GORM | 206,664 | 79,744 | 2,116 |
| SQLX | 364,005 | 52,194 ⭐ | 2,232 |
| XORM | 494,382 | 107,849 | 2,750 |

#### GetByID（按主键查询单条记录）

| ORM | ns/op | B/op | allocs/op |
|-----|-------|------|-----------|
| BORM | 4,943 ⭐ | 1,139 ⭐ | 58 |
| ZORM | 5,061 | 1,139 | 58 |
| GORM | 8,618 | 4,254 | 98 |
| SQLX | 9,508 | 1,355 ⭐ | 62 |
| XORM | 15,006 | 5,095 | 173 |

#### GetByIDs（按主键批量查询 - 10 条记录）

| ORM | ns/op | B/op | allocs/op |
|-----|-------|------|-----------|
| ZORM | 19,725 ⭐ | 4,583 ⭐ | 229 |
| BORM | 19,728 ⭐ | 4,583 ⭐ | 229 |
| SQLX | 25,820 | 5,123 | 250 |
| GORM | 27,167 | 8,419 | 320 |
| XORM | 40,291 | 14,048 | 542 |

#### Update（记录更新）

| ORM | ns/op | B/op | allocs/op |
|-----|-------|------|-----------|
| ZORM | 2,172 ⭐ | 510 ⭐ | 20 |
| BORM | 2,196 ⭐ | 510 ⭐ | 20 |
| GORM | 13,500 | 7,421 | 124 |
| XORM | 164,220 | 4,173 | 113 |
| SQLX | 171,572 | 722 ⭐ | 25 |

#### Delete（记录删除）

| ORM | ns/op | B/op | allocs/op |
|-----|-------|------|-----------|
| ZORM | 1,727 ⭐ | 191 ⭐ | 11 |
| BORM | 1,789 ⭐ | 191 ⭐ | 11 |
| GORM | 10,860 | 5,821 | 97 |
| XORM | 157,872 | 3,114 | 86 |
| SQLX | 179,034 | 303 ⭐ | 16 |

#### Count（统计查询）

| ORM | ns/op | B/op | allocs/op |
|-----|-------|------|-----------|
| ZORM | 1,702 ⭐ | 496 ⭐ | 23 |
| BORM | 1,717 ⭐ | 496 ⭐ | 23 |
| GORM | 4,473 | 2,792 | 44 |
| SQLX | 6,783 | 568 ⭐ | 26 |
| XORM | 9,188 | 2,457 | 71 |

#### GetAll（分页查询 - limit 100, offset）

| ORM | ns/op | B/op | allocs/op |
|-----|-------|------|-----------|
| ZORM | 111,041 | 23,632 | 1,723 |
| BORM | 112,534 | 23,632 | 1,723 |
| SQLX | 128,347 | 26,328 | 1,828 |
| GORM | 156,564 | 34,004 | 2,242 |
| XORM | 201,267 | 82,534 | 3,893 |

**详细结果**: 查看英文版 `README.md` 中的 “Detailed Results” 小节获取完整基准测试数据。

## 测试数据库

所有基准测试使用 **SQLite** 作为测试数据库：
- 内存数据库以获得快速性能
- 测试后自动清理临时文件
- 每个 ORM 使用独立的数据库实例

## 项目结构

```
goorm/
├── gorm/          # GORM 实现
├── xorm/          # XORM 实现
├── zorm/          # ZORM 实现
├── sqlx/           # SQLX 实现
├── borm/           # BORM 实现
├── internal/
│   ├── models/     # 测试模型 (User, Post)
│   └── orm/        # 统一的 ORM 接口
├── goorm_test.go   # 基准测试
├── go.mod          # Go 模块文件
└── README.md       # 本文件
```

## 添加新的 ORM 库

要添加新的 ORM 库：

1. 创建新目录（例如 `ent/`）
2. 在新文件中实现 `ORMInterface`
3. 在 `goorm_test.go` 的 `orms` map 中添加 ORM
4. 按照命名模式添加基准测试函数

## 贡献

欢迎贡献！请随时提交 Pull Request。

## 许可证

MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 参考

- 灵感来自 [benchplus/gocache](https://github.com/benchplus/gocache)
- [GORM 文档](https://gorm.io/docs/)
- [XORM 文档](https://xorm.io/docs/)
- [ZORM 仓库](https://github.com/IceWhaleTech/zorm)
- [SQLX 文档](https://jmoiron.github.io/sqlx/)
- [BORM 仓库](https://github.com/orca-zhang/borm)
