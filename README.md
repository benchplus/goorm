# goorm

Continuous Benchmark for ORM libraries written in golang.

[English](README.md) | [中文](README_CN.md)

## Overview

This repository provides performance benchmarks for popular Go ORM libraries using SQLite as the test database. The benchmarks are designed to compare the performance of different ORM libraries across common database operations.

## Tested ORM Libraries

- **[ZORM](https://github.com/IceWhaleTech/zorm)** - A simple, ultra-fast ORM library
- **[BORM](https://github.com/orca-zhang/borm)** - A better ORM library that is simple, fast and self-mockable for Go
- **[GORM](https://gorm.io/)** - The most popular ORM library for Go
- **[SQLX](https://github.com/jmoiron/sqlx)** - A library which provides a set of extensions on go's standard `database/sql` library
- **[XORM](https://xorm.io/)** - A simple and powerful ORM for Go

## Benchmark Tests

The following operations are benchmarked:

| Test Case | Description |
|-----------|-------------|
| `InsertSingle` | Single record insertion performance |
| `InsertBatch` | Batch insertion performance (100 records per batch) |
| `GetByID` | Single record retrieval by primary key |
| `GetByIDs` | Multiple records retrieval by primary keys |
| `Update` | Record update performance |
| `Delete` | Record deletion performance |
| `Count` | Count query performance |
| `GetAll` | Paginated query performance (limit/offset) |

## Running Benchmarks

### Prerequisites

- Go 1.21 or higher
- SQLite3

### Run All Benchmarks

```bash
go test -bench=. -benchmem
```

### Run Specific ORM Benchmark

```bash
# ZORM only
go test -bench=ZORM -benchmem

# BORM only
go test -bench=BORM -benchmem

# GORM only
go test -bench=GORM -benchmem

# XORM only
go test -bench=XORM -benchmem

# SQLX only
go test -bench=SQLX -benchmem
```

### Run Specific Test Case

```bash
# Insert single test for all ORMs
go test -bench=InsertSingle -benchmem

# GetByID test for all ORMs
go test -bench=GetByID -benchmem
```

## Benchmark Results

The benchmark results show:
- **ns/op**: Nanoseconds per operation
- **B/op**: Bytes allocated per operation
- **allocs/op**: Number of allocations per operation

### Quick Summary

| Test Case | [**ZORM**](https://github.com/IceWhaleTech/zorm) | [**BORM**](https://github.com/orca-zhang/borm) | [**GORM**](https://gorm.io/) | [**SQLX**](https://github.com/jmoiron/sqlx) | [**XORM**](https://xorm.io/) |
|-----------|---------------------------------------------------|-------------------------------------------------|------------------------------|---------------------------------------------|------------------------------|
| InsertSingle | 🚀 | 🚀 | ✈️ | 🐌 | 🐌 |
| InsertBatch | 🚀 | 🚀 | ✈️ | 🐌 | 🐌 |
| GetByID | 🚀 | 🚀 | ✈️ | ✈️ | 🐌 |
| GetByIDs | 🚀 | 🚀 | ✈️ | ✈️ | 🐌 |
| Update | 🚀 | 🚀 | ✈️ | 🐌 | 🐌 |
| Delete | 🚀 | 🚀 | ✈️ | 🐌 | 🐌 |
| Count | 🚀 | 🚀 | ✈️ | 🐌 | 🐌 |
| GetAll | 🚀 | 🚀 | ✈️ | ✈️ | 🐌 |

> 🐌 for very-slow, ✈️ for fast, 🚀 for very-fast.
>
> ⭐ indicates the ORM is **both fast and memory-efficient** for this test case (Pareto-optimal in **ns/op** and **B/op**, lower is better). Stars are placed in the **ns/op** and **B/op** columns.

### Detailed Results

#### Test Environment

- **Go Version**: 1.21+
- **Database**: SQLite (in-memory, DSN uses `cache=shared&mode=memory`)
- **CPU**: M4 Pro
- **OS**: darwin (amd64)

#### InsertSingle (Single Record Insertion)

<table>
<thead>
<tr>
<th>ORM</th>
<th>ns/op</th>
<th>Ratio</th>
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

#### InsertBatch (Batch Insertion - 100 records)

<table>
<thead>
<tr>
<th>ORM</th>
<th>ns/op</th>
<th>Ratio</th>
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

#### GetByID (Single Record Retrieval)

<table>
<thead>
<tr>
<th>ORM</th>
<th>ns/op</th>
<th>Ratio</th>
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

#### GetByIDs (Multiple Records Retrieval - 10 records)

<table>
<thead>
<tr>
<th>ORM</th>
<th>ns/op</th>
<th>Ratio</th>
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

#### Update (Record Update)

<table>
<thead>
<tr>
<th>ORM</th>
<th>ns/op</th>
<th>Ratio</th>
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

#### Delete (Record Deletion)

<table>
<thead>
<tr>
<th>ORM</th>
<th>ns/op</th>
<th>Ratio</th>
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

#### Count (Count Query)

<table>
<thead>
<tr>
<th>ORM</th>
<th>ns/op</th>
<th>Ratio</th>
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

#### GetAll (Paginated Query - limit 100, offset)

<table>
<thead>
<tr>
<th>ORM</th>
<th>ns/op</th>
<th>Ratio</th>
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

## Test Database

All benchmarks use **SQLite** as the test database:
- In-memory database for fast performance
- Temporary files are automatically cleaned up after tests
- Each ORM uses its own isolated database instance

## Project Structure

```
goorm/
├── gorm/          # GORM implementation
├── xorm/          # XORM implementation
├── zorm/          # ZORM implementation
├── sqlx/           # SQLX implementation
├── borm/           # BORM implementation
├── internal/
│   ├── models/     # Test models (User, Post)
│   └── orm/        # Unified ORM interface
├── goorm_test.go   # Benchmark tests
├── go.mod          # Go module file
└── README.md       # This file
```

## Adding New ORM Libraries

To add a new ORM library:

1. Create a new directory (e.g., `ent/`)
2. Implement the `ORMInterface` in a new file
3. Add the ORM to the `orms` map in `goorm_test.go`
4. Add benchmark functions following the naming pattern

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see [LICENSE](LICENSE) file for details.

## References

- Inspired by [benchplus/gocache](https://github.com/benchplus/gocache)
- [ZORM Repository](https://github.com/IceWhaleTech/zorm)
- [BORM Repository](https://github.com/orca-zhang/borm)
- [GORM Documentation](https://gorm.io/docs/)
- [XORM Documentation](https://xorm.io/docs/)
- [SQLX Documentation](https://jmoiron.github.io/sqlx/)
