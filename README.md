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

| ORM | ns/op | B/op | allocs/op |
|-----|-------|------|-----------|
| BORM | 2,970 ⭐ | 512 ⭐ | 22 |
| ZORM | 3,439 ⭐ | 512 ⭐ | 22 |
| GORM | 19,041 | 6,384 | 125 |
| SQLX | 183,124 | 692 ⭐ | 26 |
| XORM | 193,387 | 2,749 | 62 |

#### InsertBatch (Batch Insertion - 100 records)

| ORM | ns/op | B/op | allocs/op |
|-----|-------|------|-----------|
| BORM | 101,798 ⭐ | 61,914 ⭐ | 1,216 |
| ZORM | 103,583 ⭐ | 61,913 ⭐ | 1,216 |
| GORM | 206,664 | 79,744 | 2,116 |
| SQLX | 364,005 | 52,194 ⭐ | 2,232 |
| XORM | 494,382 | 107,849 | 2,750 |

#### GetByID (Single Record Retrieval)

| ORM | ns/op | B/op | allocs/op |
|-----|-------|------|-----------|
| BORM | 4,943 ⭐ | 1,139 ⭐ | 58 |
| ZORM | 5,061 | 1,139 | 58 |
| GORM | 8,618 | 4,254 | 98 |
| SQLX | 9,508 | 1,355 ⭐ | 62 |
| XORM | 15,006 | 5,095 | 173 |

#### GetByIDs (Multiple Records Retrieval - 10 records)

| ORM | ns/op | B/op | allocs/op |
|-----|-------|------|-----------|
| ZORM | 19,725 ⭐ | 4,583 ⭐ | 229 |
| BORM | 19,728 ⭐ | 4,583 ⭐ | 229 |
| SQLX | 25,820 | 5,123 | 250 |
| GORM | 27,167 | 8,419 | 320 |
| XORM | 40,291 | 14,048 | 542 |

#### Update (Record Update)

| ORM | ns/op | B/op | allocs/op |
|-----|-------|------|-----------|
| ZORM | 2,172 ⭐ | 510 ⭐ | 20 |
| BORM | 2,196 ⭐ | 510 ⭐ | 20 |
| GORM | 13,500 | 7,421 | 124 |
| XORM | 164,220 | 4,173 | 113 |
| SQLX | 171,572 | 722 ⭐ | 25 |

#### Delete (Record Deletion)

| ORM | ns/op | B/op | allocs/op |
|-----|-------|------|-----------|
| ZORM | 1,727 ⭐ | 191 ⭐ | 11 |
| BORM | 1,789 ⭐ | 191 ⭐ | 11 |
| GORM | 10,860 | 5,821 | 97 |
| XORM | 157,872 | 3,114 | 86 |
| SQLX | 179,034 | 303 ⭐ | 16 |

#### Count (Count Query)

| ORM | ns/op | B/op | allocs/op |
|-----|-------|------|-----------|
| ZORM | 1,702 ⭐ | 496 ⭐ | 23 |
| BORM | 1,717 ⭐ | 496 ⭐ | 23 |
| GORM | 4,473 | 2,792 | 44 |
| SQLX | 6,783 | 568 ⭐ | 26 |
| XORM | 9,188 | 2,457 | 71 |

#### GetAll (Paginated Query - limit 100, offset)

| ORM | ns/op | B/op | allocs/op |
|-----|-------|------|-----------|
| ZORM | 111,041 | 23,632 | 1,723 |
| BORM | 112,534 | 23,632 | 1,723 |
| SQLX | 128,347 | 26,328 | 1,828 |
| GORM | 156,564 | 34,004 | 2,242 |
| XORM | 201,267 | 82,534 | 3,893 |

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
- [GORM Documentation](https://gorm.io/docs/)
- [XORM Documentation](https://xorm.io/docs/)
- [ZORM Repository](https://github.com/IceWhaleTech/zorm)
- [SQLX Documentation](https://jmoiron.github.io/sqlx/)
- [BORM Repository](https://github.com/orca-zhang/borm)
