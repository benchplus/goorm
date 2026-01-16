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
- **[ENT](https://github.com/ent/ent)** - An entity framework for Go
- **[BUN](https://github.com/uptrace/bun)** - SQL-first Golang ORM

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

# BUN only
go test -bench=BUN -benchmem

# ENT only
go test -bench=ENT -benchmem
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

| Test Case | [**ZORM**](https://github.com/IceWhaleTech/zorm) | [**BORM**](https://github.com/orca-zhang/borm) | [**BUN**](https://bun.uptrace.dev/) | [**ENT**](https://github.com/ent/ent) | [**GORM**](https://gorm.io/) | [**SQLX**](https://github.com/jmoiron/sqlx) | [**XORM**](https://xorm.io/) |
|-----------|---------------------------------------------------|-------------------------------------------------|-------------------------------------|----------------------------------------|------------------------------|---------------------------------------------|------------------------------|
| InsertSingle | 🚀 | 🚀 | ✈️ | ✈️ | ✈️ | 🐌 | 🐌 |
| InsertBatch | 🚀 | 🚀 | ✈️ | 🐌 | ✈️ | 🐌 | 🐌 |
| GetByID | 🚀 | 🚀 | ✈️ | ✈️ | ✈️ | ✈️ | 🐌 |
| GetByIDs | 🚀 | 🚀 | ✈️ | ✈️ | ✈️ | ✈️ | 🐌 |
| Update | 🚀 | 🚀 | ✈️ | ✈️ | ✈️ | 🐌 | 🐌 |
| Delete | 🚀 | 🚀 | ✈️ | ✈️ | ✈️ | 🐌 | 🐌 |
| Count | 🚀 | 🚀 | ✈️ | 🐌 | ✈️ | ✈️ | 🐌 |
| GetAll | 🚀 | 🚀 | ✈️ | ✈️ | ✈️ | ✈️ | 🐌 |

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
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>2,913 ⭐</td><td>🟢 1.00x</td><td>464 ⭐</td><td>16</td></tr>
<tr style="background-color: #4CAF50;"><td>BORM</td><td>2,916</td><td>🟢 1.00x</td><td>464</td><td>16</td></tr>
<tr style="background-color: #FFA500;"><td>BUN</td><td>9,123</td><td>🟠 3.13x</td><td>5,405</td><td>27</td></tr>
<tr style="background-color: #FFA500;"><td>ENT</td><td>10,065</td><td>🟠 3.46x</td><td>2,642</td><td>74</td></tr>
<tr style="background-color: #FF6347;"><td>GORM</td><td>20,652</td><td>🔴 7.09x</td><td>6,116</td><td>96</td></tr>
<tr style="background-color: #FF6347;"><td>SQLX</td><td>176,553</td><td>🔴 60.61x</td><td>632</td><td>18</td></tr>
<tr style="background-color: #FF6347;"><td>XORM</td><td>178,033</td><td>🔴 61.12x</td><td>2,690</td><td>54</td></tr>
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
<tr style="background-color: #4CAF50;"><td>BORM</td><td>104,405</td><td>🟢 1.00x</td><td>59,503</td><td>912</td></tr>
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>104,862</td><td>🟢 1.00x</td><td>59,502 ⭐</td><td>912</td></tr>
<tr style="background-color: #FFC107;"><td>BUN</td><td>136,110</td><td>🟡 1.30x</td><td>24,455</td><td>723</td></tr>
<tr style="background-color: #FFC107;"><td>GORM</td><td>197,739</td><td>🟡 1.89x</td><td>74,929</td><td>1,494</td></tr>
<tr style="background-color: #FFA500;"><td>ENT</td><td>261,124</td><td>🟠 2.50x</td><td>213,805</td><td>3,360</td></tr>
<tr style="background-color: #FFA500;"><td>XORM</td><td>347,846</td><td>🟠 3.33x</td><td>98,830</td><td>2,449</td></tr>
<tr style="background-color: #FFA500;"><td>SQLX</td><td>373,004</td><td>🟠 3.57x</td><td>47,349</td><td>1,622</td></tr>
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
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>4,729 ⭐</td><td>🟢 1.00x</td><td>939 ⭐</td><td>33</td></tr>
<tr style="background-color: #4CAF50;"><td>BORM</td><td>4,776</td><td>🟢 1.01x</td><td>939</td><td>33</td></tr>
<tr style="background-color: #FFC107;"><td>BUN</td><td>7,191</td><td>🟡 1.52x</td><td>5,700</td><td>36</td></tr>
<tr style="background-color: #FFC107;"><td>ENT</td><td>8,766</td><td>🟡 1.85x</td><td>3,812</td><td>103</td></tr>
<tr style="background-color: #FFC107;"><td>GORM</td><td>8,975</td><td>🟡 1.90x</td><td>4,076</td><td>73</td></tr>
<tr style="background-color: #FFA500;"><td>SQLX</td><td>9,471</td><td>🟠 2.00x</td><td>1,155</td><td>37</td></tr>
<tr style="background-color: #FFA500;"><td>XORM</td><td>14,750</td><td>🟠 3.12x</td><td>4,809</td><td>139</td></tr>
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
<tr style="background-color: #4CAF50;"><td>BORM</td><td>18,082 ⭐</td><td>🟢 1.00x</td><td>3,511 ⭐</td><td>95</td></tr>
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>18,257</td><td>🟢 1.01x</td><td>3,511</td><td>95</td></tr>
<tr style="background-color: #FFC107;"><td>BUN</td><td>21,171</td><td>🟢 1.17x</td><td>7,467</td><td>107</td></tr>
<tr style="background-color: #FFC107;"><td>SQLX</td><td>24,623</td><td>🟡 1.36x</td><td>4,051</td><td>116</td></tr>
<tr style="background-color: #FFC107;"><td>ENT</td><td>24,997</td><td>🟡 1.38x</td><td>9,679</td><td>230</td></tr>
<tr style="background-color: #FFC107;"><td>GORM</td><td>25,137</td><td>🟡 1.39x</td><td>7,369</td><td>186</td></tr>
<tr style="background-color: #FFC107;"><td>XORM</td><td>35,883</td><td>🟡 1.98x</td><td>12,918</td><td>400</td></tr>
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
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>2,103 ⭐</td><td>🟢 1.00x</td><td>454 ⭐</td><td>13</td></tr>
<tr style="background-color: #4CAF50;"><td>BORM</td><td>2,107</td><td>🟢 1.00x</td><td>454</td><td>13</td></tr>
<tr style="background-color: #FFA500;"><td>BUN</td><td>5,607</td><td>🟠 2.67x</td><td>5,044</td><td>15</td></tr>
<tr style="background-color: #FF6347;"><td>GORM</td><td>14,845</td><td>🔴 7.06x</td><td>7,442</td><td>99</td></tr>
<tr style="background-color: #FF6347;"><td>ENT</td><td>20,745</td><td>🔴 9.86x</td><td>5,608</td><td>156</td></tr>
<tr style="background-color: #FF6347;"><td>SQLX</td><td>173,533</td><td>🔴 82.52x</td><td>654</td><td>16</td></tr>
<tr style="background-color: #FF6347;"><td>XORM</td><td>176,658</td><td>🔴 84.00x</td><td>4,082</td><td>103</td></tr>
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
<tr style="background-color: #4CAF50;"><td>BORM</td><td>1,646</td><td>🟢 1.00x</td><td>160</td><td>7</td></tr>
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>1,667</td><td>🟢 1.01x</td><td>159 ⭐</td><td>7</td></tr>
<tr style="background-color: #FFA500;"><td>BUN</td><td>3,798</td><td>🟠 2.31x</td><td>4,880</td><td>14</td></tr>
<tr style="background-color: #FFA500;"><td>ENT</td><td>4,318</td><td>🟠 2.62x</td><td>1,832</td><td>44</td></tr>
<tr style="background-color: #FF6347;"><td>GORM</td><td>10,536</td><td>🔴 6.40x</td><td>5,571</td><td>75</td></tr>
<tr style="background-color: #FF6347;"><td>XORM</td><td>167,642</td><td>🔴 101.85x</td><td>3,043</td><td>80</td></tr>
<tr style="background-color: #FF6347;"><td>SQLX</td><td>174,217</td><td>🔴 105.84x</td><td>255</td><td>10</td></tr>
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
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>1,546 ⭐</td><td>🟢 1.00x</td><td>440 ⭐</td><td>14</td></tr>
<tr style="background-color: #4CAF50;"><td>BORM</td><td>1,555</td><td>🟢 1.01x</td><td>440</td><td>14</td></tr>
<tr style="background-color: #FFA500;"><td>BUN</td><td>3,327</td><td>🟠 2.15x</td><td>1,288</td><td>23</td></tr>
<tr style="background-color: #FFA500;"><td>GORM</td><td>4,617</td><td>🟠 2.99x</td><td>2,720</td><td>33</td></tr>
<tr style="background-color: #FFA500;"><td>SQLX</td><td>6,716</td><td>🟠 4.34x</td><td>504</td><td>16</td></tr>
<tr style="background-color: #FF6347;"><td>XORM</td><td>9,193</td><td>🔴 5.95x</td><td>2,394</td><td>61</td></tr>
<tr style="background-color: #FF6347;"><td>ENT</td><td>20,720</td><td>🔴 13.40x</td><td>2,384</td><td>54</td></tr>
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
<tr style="background-color: #4CAF50;"><td>ZORM</td><td>95,365 ⭐</td><td>🟢 1.00x</td><td>14,829 ⭐</td><td>607</td></tr>
<tr style="background-color: #4CAF50;"><td>BORM</td><td>95,795</td><td>🟢 1.00x</td><td>14,829</td><td>607</td></tr>
<tr style="background-color: #FFC107;"><td>BUN</td><td>109,081</td><td>🟢 1.14x</td><td>21,955</td><td>709</td></tr>
<tr style="background-color: #FFC107;"><td>SQLX</td><td>112,556</td><td>🟢 1.18x</td><td>17,526</td><td>712</td></tr>
<tr style="background-color: #FFC107;"><td>ENT</td><td>115,475</td><td>🟡 1.21x</td><td>43,301</td><td>1,256</td></tr>
<tr style="background-color: #FFC107;"><td>GORM</td><td>136,701</td><td>🟡 1.43x</td><td>25,311</td><td>1,128</td></tr>
<tr style="background-color: #FFC107;"><td>XORM</td><td>182,305</td><td>🟡 1.91x</td><td>74,869</td><td>2,771</td></tr>
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
├── bun/            # BUN implementation
├── ent/             # ENT implementation
│   └── schema/      # ENT schema definitions
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

**Note for ENT**: ENT requires code generation. After adding the schema files, run:
```bash
go generate ./ent
```
This will generate the ENT client code needed for the implementation.

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
- [ENT Documentation](https://entgo.io/)
- [BUN Documentation](https://bun.uptrace.dev/)