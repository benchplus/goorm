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
# 仅 ZORM
go test -bench=ZORM -benchmem

# 仅 BORM
go test -bench=BORM -benchmem

# 仅 GORM
go test -bench=GORM -benchmem

# 仅 XORM
go test -bench=XORM -benchmem

# 仅 SQLX
go test -bench=SQLX -benchmem
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

<table>
<thead>
<tr>
<th>ORM</th>
<th>ns/op</th>
<th>倍数</th>
<th>B/op</th>
<th>allocs/op</th>
</tr>
</thead>
<tbody>
<tr style="background-color: #4CAF50;"><td>BORM</td><td>2,970 ⭐</td><td>1.00x</td><td>512 ⭐</td><td>22</td></tr>
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>3,439 ⭐</td><td>1.16x</td><td>512 ⭐</td><td>22</td></tr>
<tr style="background-color: #FFA500;"><td>GORM</td><td>19,041</td><td>6.41x</td><td>6,384</td><td>125</td></tr>
<tr style="background-color: #FF6347;"><td>SQLX</td><td>183,124</td><td>61.66x</td><td>692 ⭐</td><td>26</td></tr>
<tr style="background-color: #FF6347;"><td>XORM</td><td>193,387</td><td>65.11x</td><td>2,749</td><td>62</td></tr>
</tbody>
</table>

#### InsertBatch（批量插入 - 每批 100 条）

<table>
<thead>
<tr>
<th>ORM</th>
<th>ns/op</th>
<th>倍数</th>
<th>B/op</th>
<th>allocs/op</th>
</tr>
</thead>
<tbody>
<tr style="background-color: #4CAF50;"><td>BORM</td><td>101,798 ⭐</td><td>1.00x</td><td>61,914 ⭐</td><td>1,216</td></tr>
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>103,583 ⭐</td><td>1.02x</td><td>61,913 ⭐</td><td>1,216</td></tr>
<tr style="background-color: #FFC107;"><td>GORM</td><td>206,664</td><td>2.03x</td><td>79,744</td><td>2,116</td></tr>
<tr style="background-color: #FFA500;"><td>SQLX</td><td>364,005</td><td>3.58x</td><td>52,194 ⭐</td><td>2,232</td></tr>
<tr style="background-color: #FFA500;"><td>XORM</td><td>494,382</td><td>4.86x</td><td>107,849</td><td>2,750</td></tr>
</tbody>
</table>

#### GetByID（按主键查询单条记录）

<table>
<thead>
<tr>
<th>ORM</th>
<th>ns/op</th>
<th>倍数</th>
<th>B/op</th>
<th>allocs/op</th>
</tr>
</thead>
<tbody>
<tr style="background-color: #4CAF50;"><td>BORM</td><td>4,943 ⭐</td><td>1.00x</td><td>1,139 ⭐</td><td>58</td></tr>
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>5,061 ⭐</td><td>1.02x</td><td>1,139 ⭐</td><td>58</td></tr>
<tr style="background-color: #FFC107;"><td>GORM</td><td>8,618</td><td>1.74x</td><td>4,254</td><td>98</td></tr>
<tr style="background-color: #FFC107;"><td>SQLX</td><td>9,508</td><td>1.92x</td><td>1,355 ⭐</td><td>62</td></tr>
<tr style="background-color: #FFA500;"><td>XORM</td><td>15,006</td><td>3.03x</td><td>5,095</td><td>173</td></tr>
</tbody>
</table>

#### GetByIDs（按主键批量查询 - 10 条记录）

<table>
<thead>
<tr>
<th>ORM</th>
<th>ns/op</th>
<th>倍数</th>
<th>B/op</th>
<th>allocs/op</th>
</tr>
</thead>
<tbody>
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>19,725 ⭐</td><td>1.00x</td><td>4,583 ⭐</td><td>229</td></tr>
<tr style="background-color: #4CAF50;"><td>BORM</td><td>19,728 ⭐</td><td>1.00x</td><td>4,583 ⭐</td><td>229</td></tr>
<tr style="background-color: #FFC107;"><td>SQLX</td><td>25,820</td><td>1.31x</td><td>5,123</td><td>250</td></tr>
<tr style="background-color: #FFC107;"><td>GORM</td><td>27,167</td><td>1.38x</td><td>8,419</td><td>320</td></tr>
<tr style="background-color: #FFA500;"><td>XORM</td><td>40,291</td><td>2.04x</td><td>14,048</td><td>542</td></tr>
</tbody>
</table>

#### Update（记录更新）

<table>
<thead>
<tr>
<th>ORM</th>
<th>ns/op</th>
<th>倍数</th>
<th>B/op</th>
<th>allocs/op</th>
</tr>
</thead>
<tbody>
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>2,172 ⭐</td><td>1.00x</td><td>510 ⭐</td><td>20</td></tr>
<tr style="background-color: #4CAF50;"><td>BORM</td><td>2,196 ⭐</td><td>1.01x</td><td>510 ⭐</td><td>20</td></tr>
<tr style="background-color: #FFA500;"><td>GORM</td><td>13,500</td><td>6.22x</td><td>7,421</td><td>124</td></tr>
<tr style="background-color: #FF6347;"><td>XORM</td><td>164,220</td><td>75.66x</td><td>4,173</td><td>113</td></tr>
<tr style="background-color: #FF6347;"><td>SQLX</td><td>171,572</td><td>79.01x</td><td>722 ⭐</td><td>25</td></tr>
</tbody>
</table>

#### Delete（记录删除）

<table>
<thead>
<tr>
<th>ORM</th>
<th>ns/op</th>
<th>倍数</th>
<th>B/op</th>
<th>allocs/op</th>
</tr>
</thead>
<tbody>
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>1,727 ⭐</td><td>1.00x</td><td>191 ⭐</td><td>11</td></tr>
<tr style="background-color: #4CAF50;"><td>BORM</td><td>1,789 ⭐</td><td>1.04x</td><td>191 ⭐</td><td>11</td></tr>
<tr style="background-color: #FFA500;"><td>GORM</td><td>10,860</td><td>6.29x</td><td>5,821</td><td>97</td></tr>
<tr style="background-color: #FF6347;"><td>XORM</td><td>157,872</td><td>91.41x</td><td>3,114</td><td>86</td></tr>
<tr style="background-color: #FF6347;"><td>SQLX</td><td>179,034</td><td>103.67x</td><td>303 ⭐</td><td>16</td></tr>
</tbody>
</table>

#### Count（统计查询）

<table>
<thead>
<tr>
<th>ORM</th>
<th>ns/op</th>
<th>倍数</th>
<th>B/op</th>
<th>allocs/op</th>
</tr>
</thead>
<tbody>
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>1,702 ⭐</td><td>1.00x</td><td>496 ⭐</td><td>23</td></tr>
<tr style="background-color: #4CAF50;"><td>BORM</td><td>1,717 ⭐</td><td>1.01x</td><td>496 ⭐</td><td>23</td></tr>
<tr style="background-color: #FFA500;"><td>GORM</td><td>4,473</td><td>2.63x</td><td>2,792</td><td>44</td></tr>
<tr style="background-color: #FFA500;"><td>SQLX</td><td>6,783</td><td>3.98x</td><td>568 ⭐</td><td>26</td></tr>
<tr style="background-color: #FFA500;"><td>XORM</td><td>9,188</td><td>5.40x</td><td>2,457</td><td>71</td></tr>
</tbody>
</table>

#### GetAll（分页查询 - limit 100, offset）

<table>
<thead>
<tr>
<th>ORM</th>
<th>ns/op</th>
<th>倍数</th>
<th>B/op</th>
<th>allocs/op</th>
</tr>
</thead>
<tbody>
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>111,041</td><td>1.00x</td><td>23,632</td><td>1,723</td></tr>
<tr style="background-color: #4CAF50;"><td>BORM</td><td>112,534</td><td>1.01x</td><td>23,632</td><td>1,723</td></tr>
<tr style="background-color: #FFC107;"><td>SQLX</td><td>128,347</td><td>1.16x</td><td>26,328</td><td>1,828</td></tr>
<tr style="background-color: #FFC107;"><td>GORM</td><td>156,564</td><td>1.41x</td><td>34,004</td><td>2,242</td></tr>
<tr style="background-color: #FFC107;"><td>XORM</td><td>201,267</td><td>1.81x</td><td>82,534</td><td>3,893</td></tr>
</tbody>
</table>

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
- [ZORM 文档](https://github.com/IceWhaleTech/zorm)
- [BORM 文档](https://github.com/orca-zhang/borm)
- [GORM 文档](https://gorm.io/docs/)
- [XORM 文档](https://xorm.io/docs/)
- [SQLX 文档](https://jmoiron.github.io/sqlx/)
