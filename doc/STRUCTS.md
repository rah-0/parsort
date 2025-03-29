# Struct Sorting

Moving away from **multiway parallel merge sort with pairwise merging**, the algorithm used in this case is **Parallel Stable Pairwise Merge Sort**.

### Why Multiway Merge Sort (with a priority queue) is not efficient for structs:

1. High Per-Comparison Cost
    - Multiway merge uses a **min-heap** (priority queue) to always extract the smallest among `k` sorted sequences.
    - For basic types like `int` or `float64`, comparisons are very cheap.
    - For structs, each comparison calls a user-defined function like `less(a, b)`, which is often **much slower** and may include:
      - Field access
      - Pointer chasing
      - Possibly nested comparisons
    - Heaps require **log k** comparisons per pop → high cost when each comparison is expensive.

2. Loss of Stability (unless handled manually)
    - A priority queue selects the smallest next element across `k` streams, but:
      - If two elements compare equal, the heap pops one arbitrarily.
      - For structs, where sorting is often based on fields like timestamps or IDs, stability is critical.
    - Preserving stability would require attaching original indices and modifying the comparison logic, adding overhead.

3. More Pointer Indirection
    - Priority queues for multiway merges usually hold items like `(value, index, chunkID)`.
    - For basic types, this is fast due to small size and value semantics.
    - For structs, this means more indirection and memory operations to track values and metadata.
    - In high-volume merges, this results in additional allocations and garbage collection pressure.

4. Cache Locality Penalty
    - Multiway merging involves simultaneous access to `k` sequences.
    - This breaks sequential access patterns and thrashes the cache.
    - For small types (e.g., `int`), this impact is minor.
    - For larger structs with non-contiguous memory layouts, cache misses increase and degrade performance.

5. Complexity for Generalized Types
    - Priority queues require a comparison interface compatible with heap operations.
    - For structs, this adds overhead: closures, interface boxing, or function indirection.
    - Supporting `less(a, b T)` in a heap-compatible form introduces complexity.
    - Pairwise merging avoids this by working directly on `[]T` and passing the `less` function transparently.

# Benchmarking

The struct used for benchmarks is:
```
type person struct {
	Name string
	Age  int
}
```
The benchmark is a simple as:
```
sort.Slice(people, func(i, j int) bool {
    return tmp[i].Age < tmp[j].Age
})
```
Which we will use as the first result:

| Size       | ns/op     | ns/op (%) | B/op        | B/op (%) |
|------------|-----------|-----------|-------------|-----------|
| 1000       | 52.13µs   |    100.00 | 24.09KiB    |    100.00 |
| 10000      | 654.21µs  |    100.00 | 240.09KiB   |    100.00 |
| 100000     | 6.40ms    |    100.00 | 2.29MiB     |    100.00 |
| 1000000    | 58.21ms   |    100.00 | 22.89MiB    |    100.00 |
| 10000000   | 553.95ms  |    100.00 | 228.88MiB   |    100.00 |
| 100000000  | 7.94s     |    100.00 | 2.24GiB     |    100.00 |

Let's see the performance with interfaces: 
```
type byAge []person
func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

sort.Sort(byAge(people))
```
| Size       | ns/op     | ns/op (%) | B/op        | B/op (%) |
|------------|-----------|-----------|-------------|----------|
| 1000       | 28.66µs   |    -45.02 | 24.02KiB    |     -0.29 |
| 10000      | 467.26µs  |    -28.58 | 240.02KiB   |     -0.03 |
| 100000     | 5.01ms    |    -21.69 | 2.29MiB     |     0.00 |
| 1000000    | 43.74ms   |    -24.86 | 22.89MiB    |     0.00 |
| 10000000   | 416.89ms  |    -24.74 | 228.88MiB   |     0.00 |
| 100000000  | 4.30s     |    -45.88 | 2.24GiB     |     0.00 |

This is already a major improvement.

So the next comparison will be relative to this, Parallel Stable Pairwise Merge Sort:

```
StructAsc(people, func(a, b person) bool {
  return a.Age < b.Age
})
```

| Size       | ns/op     | ns/op (%) | B/op          | B/op (%) |
|------------|-----------|-----------|---------------|-----------|
| 1000       | 111.50µs  |   +389.03 | 50.55KiB      |   +210.41 |
| 10000      | 679.97µs  |   +145.52 | 482.55KiB     |   +201.04 |
| 100000     | 4.03ms    |    -19.53 | 4.58MiB       |   +200.11 |
| 1000000    | 32.13ms   |    -26.55 | 45.78MiB      |   +200.01 |
| 10000000   | 249.71ms  |    -40.10 | 457.77MiB     |   +200.00 |
| 100000000  | 2.65s     |    -38.35 | 4.47GiB       |   +200.00 |

# 📊 Conclusions

- **Under 10k elements**  
  The standard library’s `sort.Sort` outperforms all alternatives. The overhead of parallelism and memory allocations isn’t worth it for small slices.

- **100k elements and above**  
  The **Parallel Stable Pairwise Merge Sort** consistently outperforms both `sort.Slice` and `sort.Sort` in raw CPU performance (`ns/op`), achieving:
   - ~20–40% faster execution
   - ~2× higher memory usage (due to parallel merge buffers)

- **Memory Tradeoff**  
  While `B/op` is around **200%** of the baseline, this is a deliberate tradeoff to maximize **CPU throughput**, which is acceptable in modern systems for performance-critical workloads.

- **Stability**  
  This parallel implementation is **fully stable**, maintaining the original order of equal elements. In contrast, `sort.Slice` and `sort.Sort` are not stable unless you use `sort.SliceStable`.

- **Best Use Cases**
   - Sorting large slices of structs (e.g., event logs, database records)
   - Multi-core systems where parallel sorting improves throughput
   - Situations requiring **stable** sorting guarantees

