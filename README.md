[![Go Report Card](https://goreportcard.com/badge/github.com/rah-0/parsort?v=1)](https://goreportcard.com/report/github.com/rah-0/parsort)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

<a href="https://www.buymeacoffee.com/rah.0" target="_blank">
  <img src="https://cdn.buymeacoffee.com/buttons/v2/arial-orange.png" alt="Buy Me A Coffee" height="50" style="height:50px;">
</a>


# (par)allel (sort)ing
A minimal, dependency-free Go library for parallel sorting of slices. Designed to scale across CPU cores and stay close to idiomatic Go.

## Installation

```bash
go get github.com/rah-0/parsort
```

## Algorithm
The algorithm built is a **parallel merge sort**, more specifically a **multiway parallel merge sort with pairwise merging**.

### Breakdown:
- **Divide phase**: The input slice is split into `N` chunks (`N = NumCPU`).
- **Parallel sort phase**: Each chunk is sorted independently using `sort.Slice` in parallel goroutines.
- **Merge phase**: All sorted chunks are merged in **parallel pairwise steps** (`log₂N steps total`).

This approach isn't a classic recursive merge sort — instead, it's:
- **Iterative**, not recursive.
- Uses **parallelism for both sorting and merging**.
- Merge is **pairwise**, not `N-way` (N-way can be faster for large chunk sets but more complex).

### Why not use generics?
1. Generics introduce performance overhead
    - Generic `sort.Slice` requires a comparison function (e.g., `func(i, j int) bool`).
    - This introduces:
      - **Function call overhead** for every comparison.
      - **Indirect branching**, which harms CPU branch prediction and cache locality.
    - Native sorts like `sort.Ints`, `sort.Float64s`, etc. are **heavily optimized** and compiled with inlined, tight loops and direct comparisons.
2. Avoiding heap allocations
   - Generic code may cause additional heap allocations (captured closures, interfaces).
   - Currently, allocate only when merging, minimizing GC pressure.
3. Use case is performance-focused
    - Generic versions are `1.5x–2x` slower than concrete ones.
    - Each type (`int`, `float64`, `string`, etc.) gets the **best possible performance**.
    - Since this a sorting library, it's worth specializing.

## Performance
Overall, ***parsort*** can reduce ns/op by up to **90%** but at the expense of around **3** times more memory.

### Int
| Size       | Order | ns/op (%) | B/op (%) |
|------------|-------|-----------|----------|
| 10000      | Asc   |   -7.09   |  +301.81 |
| 10000      | Desc  |  -59.20   |  +301.63 |
| 100000     | Asc   |  -35.48   |  +302.22 |
| 100000     | Desc  |  -70.23   |  +302.17 |
| 1000000    | Asc   |  -72.71   |  +300.43 |
| 1000000    | Desc  |  -87.76   |  +300.43 |
| 10000000   | Asc   |  -74.58   |  +300.02 |
| 10000000   | Desc  |  -88.51   |  +300.02 |

### Float
| Size       | Order | ns/op (%) | B/op (%) |
|------------|-------|-----------|----------|
| 10000      | Asc   |   -22.90  |  +301.82 |
| 10000      | Desc  |   -59.62  |  +301.63 |
| 100000     | Asc   |   -52.34  |  +302.23 |
| 100000     | Desc  |   -75.61  |  +302.20 |
| 1000000    | Asc   |   -73.67  |  +300.43 |
| 1000000    | Desc  |   -86.61  |  +300.43 |
| 10000000   | Asc   |   -76.69  |  +300.02 |
| 10000000   | Desc  |   -88.17  |  +300.02 |

### String
| Size       | Order | ns/op (%) | B/op (%) |
|------------|-------|-----------|----------|
| 10000      | Asc   |   -21.91  |  +300.92 |
| 10000      | Desc  |   -44.02  |  +300.83 |
| 100000     | Asc   |   -48.78  |  +300.09 |
| 100000     | Desc  |   -63.57  |  +300.08 |
| 1000000    | Asc   |   -52.03  |  +300.11 |
| 1000000    | Desc  |   -66.70  |  +300.08 |
| 10000000   | Asc   |   -57.81  |  +300.00 |
| 10000000   | Desc  |   -72.57  |  +300.00 |

### time.Time
| Size       | Order    | ns/op (%) | B/op (%) |
|------------|----------|-----------|----------|
| 10000      | Asc/Desc |   -35.86  |  +307.43 |
| 100000     | Asc/Desc |   -63.49  |  +301.43 |
| 1000000    | Asc/Desc |   -70.35  |  +300.08 |
| 10000000   | Asc/Desc |   -72.07  |  +300.01 |

## Supported Types

Parsort provides specialized sorting functions for each of these types:

- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- `float32`, `float64`
- `string`
- `time.Time`
- `struct` (via generics)

Each type has both ascending (`TypeAsc`) and descending (`TypeDesc`) sorting functions.

## Performance Tuning

Parsort automatically determines if a slice is large enough to benefit from parallel sorting. The default thresholds work well for most systems, but you can optimize them for your specific hardware:

```go
// Optimize thresholds for your hardware
parsort.Tune()

// Or tune with specific parameters
parsort.TuneSpecific(5, 1000, 1000, -10.0, true)
```

Thresholds can also be manually adjusted:

```go
// Manually set threshold for int sorting
parsort.IntMinParallelSize = 5000
```

## 📌 Additional Resources
- [Struct sorting details](https://github.com/rah-0/parsort/blob/master/doc/STRUCTS.md)
- [Comparison to other libraries](https://github.com/rah-0/benchmarks/tree/master/meta#sorting)
- [Raw benchmark data](https://github.com/rah-0/parsort/blob/master/doc/BENCHMARK.md)
- [Changelog](https://github.com/rah-0/parsort/blob/master/doc/CHANGELOG.md)
- [Tuner documentation](https://github.com/rah-0/parsort/blob/master/doc/TUNER.md)

## 💚 Support
Parsort was built out of love for clean and fast code. 
If it saved you time or brought value to your project, feel free to show some support. Every bit is appreciated 🙂

[![Buy Me A Coffee](https://cdn.buymeacoffee.com/buttons/default-orange.png)](https://www.buymeacoffee.com/rah.0)
