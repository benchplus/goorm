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
- **[ENT](https://github.com/ent/ent)** - Go 的实体框架
- **[BUN](https://github.com/uptrace/bun)** - SQL优先 Golang ORM
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

# 仅 BUN
go test -bench=BUN -benchmem

# 仅 ENT
go test -bench=ENT -benchmem
```

### 运行特定测试用例

```bash
# 所有 ORM 的 InsertSingle 测试
go test -bench=InsertSingle -benchmem

# 所有 ORM 的 GetByID 测试
go test -bench=GetByID -benchmem
```

### 快速摘要

<table>
<thead>
<tr>
<th>测试用例</th>
<th><a href="https://github.com/IceWhaleTech/zorm"><strong>ZORM</strong></a></th>
<th><a href="https://github.com/orca-zhang/borm"><strong>BORM</strong></a></th>
<th><a href="https://bun.uptrace.dev/"><strong>BUN</strong></a></th>
<th><a href="https://github.com/ent/ent"><strong>ENT</strong></a></th>
<th><a href="https://gorm.io/"><strong>GORM</strong></a></th>
<th><a href="https://github.com/jmoiron/sqlx"><strong>SQLX</strong></a></th>
<th><a href="https://xorm.io/"><strong>XORM</strong></a></th>
</tr>
</thead>
<tbody>
<tr><td>InsertSingle</td><td style="background-color: #4CAF50;">🟢 1.00x</td><td style="background-color: #4CAF50;">🟢 1.00x</td><td style="background-color: #FFA500;">🟠 3.13x</td><td style="background-color: #FFA500;">🟠 3.46x</td><td style="background-color: #FF6347;">🔴 7.09x</td><td style="background-color: #FF6347;">🔴 60.61x</td><td style="background-color: #FF6347;">🔴 61.12x</td></tr>
<tr><td>InsertBatch</td><td style="background-color: #4CAF50;">🟢 1.00x</td><td style="background-color: #4CAF50;">🟢 1.00x</td><td style="background-color: #FFC107;">🟡 1.30x</td><td style="background-color: #FFA500;">🟠 2.50x</td><td style="background-color: #FFC107;">🟡 1.89x</td><td style="background-color: #FFA500;">🟠 3.57x</td><td style="background-color: #FFA500;">🟠 3.33x</td></tr>
<tr><td>GetByID</td><td style="background-color: #4CAF50;">🟢 1.00x</td><td style="background-color: #4CAF50;">🟢 1.01x</td><td style="background-color: #FFC107;">🟡 1.52x</td><td style="background-color: #FFC107;">🟡 1.85x</td><td style="background-color: #FFC107;">🟡 1.90x</td><td style="background-color: #FFA500;">🟠 2.00x</td><td style="background-color: #FFA500;">🟠 3.12x</td></tr>
<tr><td>GetByIDs</td><td style="background-color: #4CAF50;">🟢 1.01x</td><td style="background-color: #4CAF50;">🟢 1.00x</td><td style="background-color: #FFC107;">🟡 1.17x</td><td style="background-color: #FFC107;">🟡 1.38x</td><td style="background-color: #FFC107;">🟡 1.39x</td><td style="background-color: #FFC107;">🟡 1.36x</td><td style="background-color: #FFC107;">🟡 1.98x</td></tr>
<tr><td>Update</td><td style="background-color: #4CAF50;">🟢 1.00x</td><td style="background-color: #4CAF50;">🟢 1.00x</td><td style="background-color: #FFA500;">🟠 2.67x</td><td style="background-color: #FF6347;">🔴 9.86x</td><td style="background-color: #FF6347;">🔴 7.06x</td><td style="background-color: #FF6347;">🔴 82.52x</td><td style="background-color: #FF6347;">🔴 84.00x</td></tr>
<tr><td>Delete</td><td style="background-color: #4CAF50;">🟢 1.01x</td><td style="background-color: #4CAF50;">🟢 1.00x</td><td style="background-color: #FFA500;">🟠 2.31x</td><td style="background-color: #FFA500;">🟠 2.62x</td><td style="background-color: #FF6347;">🔴 6.40x</td><td style="background-color: #FF6347;">🔴 105.84x</td><td style="background-color: #FF6347;">🔴 101.85x</td></tr>
<tr><td>Count</td><td style="background-color: #4CAF50;">🟢 1.00x</td><td style="background-color: #4CAF50;">🟢 1.01x</td><td style="background-color: #FFA500;">🟠 2.15x</td><td style="background-color: #FF6347;">🔴 13.40x</td><td style="background-color: #FFA500;">🟠 2.99x</td><td style="background-color: #FFA500;">🟠 4.34x</td><td style="background-color: #FF6347;">🔴 5.95x</td></tr>
<tr><td>GetAll</td><td style="background-color: #4CAF50;">🟢 1.00x</td><td style="background-color: #4CAF50;">🟢 1.00x</td><td style="background-color: #FFC107;">🟡 1.14x</td><td style="background-color: #FFC107;">🟡 1.21x</td><td style="background-color: #FFC107;">🟡 1.43x</td><td style="background-color: #FFC107;">🟡 1.18x</td><td style="background-color: #FFC107;">🟡 1.91x</td></tr>
</tbody>
</table>

> 倍数表示相对于最快 ORM 的性能倍数（数值越小越好）
>
> ⭐ 表示该 ORM 在该测试用例下 **同时更快且更省内存**（在 **ns/op** 与 **B/op** 这两个维度上为 Pareto 最优，数值越小越好）。⭐ 标记会打在 **ns/op** 和 **B/op** 两列中。

### 详细结果

#### 测试环境

- **Go Version**: 1.21+
- **Database**: SQLite（内存数据库，DSN 使用 `cache=shared&mode=memory`）
- **CPU**: M4 Pro
- **OS**: darwin (amd64)

## 基准测试结果

基准测试结果显示：
- **ns/op**: 每次操作的纳秒数
- **B/op**: 每次操作分配的字节数
- **allocs/op**: 每次操作的分配次数

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
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>2,913 ⭐</td><td>🟢 1.00x</td><td>464 ⭐</td><td>16</td></tr>
<tr style="background-color: #4CAF50;"><td>BORM</td><td>2,916</td><td>🟢 1.00x</td><td>464</td><td>16</td></tr>
<tr style="background-color: #FFA500;"><td>BUN</td><td>9,123</td><td>🟠 3.13x</td><td>5,405</td><td>27</td></tr>
<tr style="background-color: #FFA500;"><td>ENT</td><td>10,065</td><td>🟠 3.46x</td><td>2,642</td><td>74</td></tr>
<tr style="background-color: #FF6347;"><td>GORM</td><td>20,652</td><td>🔴 7.09x</td><td>6,116</td><td>96</td></tr>
<tr style="background-color: #FF6347;"><td>SQLX</td><td>176,553</td><td>🔴 60.61x</td><td>632</td><td>18</td></tr>
<tr style="background-color: #FF6347;"><td>XORM</td><td>178,033</td><td>🔴 61.12x</td><td>2,690</td><td>54</td></tr>
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
<tr style="background-color: #4CAF50;"><td>BORM</td><td>104,405</td><td>🟢 1.00x</td><td>59,503</td><td>912</td></tr>
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>104,862</td><td>🟢 1.00x</td><td>59,502 ⭐</td><td>912</td></tr>
<tr style="background-color: #FFC107;"><td>BUN</td><td>136,110</td><td>🟡 1.30x</td><td>24,455</td><td>723</td></tr>
<tr style="background-color: #FFC107;"><td>GORM</td><td>197,739</td><td>🟡 1.89x</td><td>74,929</td><td>1,494</td></tr>
<tr style="background-color: #FFA500;"><td>ENT</td><td>261,124</td><td>🟠 2.50x</td><td>213,805</td><td>3,360</td></tr>
<tr style="background-color: #FFA500;"><td>XORM</td><td>347,846</td><td>🟠 3.33x</td><td>98,830</td><td>2,449</td></tr>
<tr style="background-color: #FFA500;"><td>SQLX</td><td>373,004</td><td>🟠 3.57x</td><td>47,349</td><td>1,622</td></tr>
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
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>4,729 ⭐</td><td>🟢 1.00x</td><td>939 ⭐</td><td>33</td></tr>
<tr style="background-color: #4CAF50;"><td>BORM</td><td>4,776</td><td>🟢 1.01x</td><td>939</td><td>33</td></tr>
<tr style="background-color: #FFC107;"><td>BUN</td><td>7,191</td><td>🟡 1.52x</td><td>5,700</td><td>36</td></tr>
<tr style="background-color: #FFC107;"><td>ENT</td><td>8,766</td><td>🟡 1.85x</td><td>3,812</td><td>103</td></tr>
<tr style="background-color: #FFC107;"><td>GORM</td><td>8,975</td><td>🟡 1.90x</td><td>4,076</td><td>73</td></tr>
<tr style="background-color: #FFA500;"><td>SQLX</td><td>9,471</td><td>🟠 2.00x</td><td>1,155</td><td>37</td></tr>
<tr style="background-color: #FFA500;"><td>XORM</td><td>14,750</td><td>🟠 3.12x</td><td>4,809</td><td>139</td></tr>
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
<tr style="background-color: #4CAF50;"><td>BORM</td><td>18,082 ⭐</td><td>🟢 1.00x</td><td>3,511 ⭐</td><td>95</td></tr>
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>18,257</td><td>🟢 1.01x</td><td>3,511</td><td>95</td></tr>
<tr style="background-color: #FFC107;"><td>BUN</td><td>21,171</td><td>🟢 1.17x</td><td>7,467</td><td>107</td></tr>
<tr style="background-color: #FFC107;"><td>SQLX</td><td>24,623</td><td>🟡 1.36x</td><td>4,051</td><td>116</td></tr>
<tr style="background-color: #FFC107;"><td>ENT</td><td>24,997</td><td>🟡 1.38x</td><td>9,679</td><td>230</td></tr>
<tr style="background-color: #FFC107;"><td>GORM</td><td>25,137</td><td>🟡 1.39x</td><td>7,369</td><td>186</td></tr>
<tr style="background-color: #FFC107;"><td>XORM</td><td>35,883</td><td>🟡 1.98x</td><td>12,918</td><td>400</td></tr>
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
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>2,103 ⭐</td><td>🟢 1.00x</td><td>454 ⭐</td><td>13</td></tr>
<tr style="background-color: #4CAF50;"><td>BORM</td><td>2,107</td><td>🟢 1.00x</td><td>454</td><td>13</td></tr>
<tr style="background-color: #FFA500;"><td>BUN</td><td>5,607</td><td>🟠 2.67x</td><td>5,044</td><td>15</td></tr>
<tr style="background-color: #FF6347;"><td>GORM</td><td>14,845</td><td>🔴 7.06x</td><td>7,442</td><td>99</td></tr>
<tr style="background-color: #FF6347;"><td>ENT</td><td>20,745</td><td>🔴 9.86x</td><td>5,608</td><td>156</td></tr>
<tr style="background-color: #FF6347;"><td>SQLX</td><td>173,533</td><td>🔴 82.52x</td><td>654</td><td>16</td></tr>
<tr style="background-color: #FF6347;"><td>XORM</td><td>176,658</td><td>🔴 84.00x</td><td>4,082</td><td>103</td></tr>
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
<tr style="background-color: #4CAF50;"><td>BORM</td><td>1,646</td><td>🟢 1.00x</td><td>160</td><td>7</td></tr>
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>1,667</td><td>🟢 1.01x</td><td>159 ⭐</td><td>7</td></tr>
<tr style="background-color: #FFA500;"><td>BUN</td><td>3,798</td><td>🟠 2.31x</td><td>4,880</td><td>14</td></tr>
<tr style="background-color: #FFA500;"><td>ENT</td><td>4,318</td><td>🟠 2.62x</td><td>1,832</td><td>44</td></tr>
<tr style="background-color: #FF6347;"><td>GORM</td><td>10,536</td><td>🔴 6.40x</td><td>5,571</td><td>75</td></tr>
<tr style="background-color: #FF6347;"><td>XORM</td><td>167,642</td><td>🔴 101.85x</td><td>3,043</td><td>80</td></tr>
<tr style="background-color: #FF6347;"><td>SQLX</td><td>174,217</td><td>🔴 105.84x</td><td>255</td><td>10</td></tr>
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
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>1,546 ⭐</td><td>🟢 1.00x</td><td>440 ⭐</td><td>14</td></tr>
<tr style="background-color: #4CAF50;"><td>BORM</td><td>1,555</td><td>🟢 1.01x</td><td>440</td><td>14</td></tr>
<tr style="background-color: #FFA500;"><td>BUN</td><td>3,327</td><td>🟠 2.15x</td><td>1,288</td><td>23</td></tr>
<tr style="background-color: #FFA500;"><td>GORM</td><td>4,617</td><td>🟠 2.99x</td><td>2,720</td><td>33</td></tr>
<tr style="background-color: #FFA500;"><td>SQLX</td><td>6,716</td><td>🟠 4.34x</td><td>504</td><td>16</td></tr>
<tr style="background-color: #FF6347;"><td>XORM</td><td>9,193</td><td>🔴 5.95x</td><td>2,394</td><td>61</td></tr>
<tr style="background-color: #FF6347;"><td>ENT</td><td>20,720</td><td>🔴 13.40x</td><td>2,384</td><td>54</td></tr>
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
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>95,365 ⭐</td><td>🟢 1.00x</td><td>14,829 ⭐</td><td>607</td></tr>
<tr style="background-color: #4CAF50;"><td>BORM</td><td>95,795</td><td>🟢 1.00x</td><td>14,829</td><td>607</td></tr>
<tr style="background-color: #FFC107;"><td>BUN</td><td>109,081</td><td>🟢 1.14x</td><td>21,955</td><td>709</td></tr>
<tr style="background-color: #FFC107;"><td>SQLX</td><td>112,556</td><td>🟢 1.18x</td><td>17,526</td><td>712</td></tr>
<tr style="background-color: #FFC107;"><td>ENT</td><td>115,475</td><td>🟡 1.21x</td><td>43,301</td><td>1,256</td></tr>
<tr style="background-color: #FFC107;"><td>GORM</td><td>136,701</td><td>🟡 1.43x</td><td>25,311</td><td>1,128</td></tr>
<tr style="background-color: #FFC107;"><td>XORM</td><td>182,305</td><td>🟡 1.91x</td><td>74,869</td><td>2,771</td></tr>
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
├── bun/            # BUN 实现
├── ent/             # ENT 实现
│   └── schema/      # ENT schema 定义
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

**ENT 注意事项**：ENT 需要代码生成。添加 schema 文件后，运行：
```bash
go generate ./ent
```
这将生成实现所需的 ENT 客户端代码。

## 贡献

欢迎贡献！请随时提交 Pull Request。

## 许可证

MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 参考

- 灵感来自 [benchplus/gocache](https://github.com/benchplus/gocache)
- [ZORM 仓库](https://github.com/IceWhaleTech/zorm)
- [BORM 仓库](https://github.com/orca-zhang/borm)
- [GORM 文档](https://gorm.io/docs/)
- [XORM 文档](https://xorm.io/docs/)
- [SQLX 文档](https://jmoiron.github.io/sqlx/)
- [ENT 文档](https://entgo.io/)
- [BUN 文档](https://bun.uptrace.dev/)