package experimental

import (
	"sort"
	"sync"
)

func MinimalCopyMergeSort[T any](data []T, less func(a, b T) bool) {
	n := len(data)
	if n <= 1 {
		return
	}

	chunkSize := (n + CoreCount - 1) / CoreCount
	chunks := make([][]T, 0, CoreCount)

	for i := 0; i < n; i += chunkSize {
		end := i + chunkSize
		if end > n {
			end = n
		}
		chunk := make([]T, end-i)
		copy(chunk, data[i:end])
		chunks = append(chunks, chunk)
	}

	var wg sync.WaitGroup
	for i := range chunks {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sort.Slice(chunks[i], func(a, b int) bool {
				return less(chunks[i][a], chunks[i][b])
			})
		}(i)
	}
	wg.Wait()

	buffer := make([]T, len(data))
	toggle := false

	// Merge chunks pairwise into shared buffer to reduce allocations
	for len(chunks) > 1 {
		mergedChunks := make([][]T, (len(chunks)+1)/2)
		var mWg sync.WaitGroup

		for i := 0; i < len(chunks); i += 2 {
			if i+1 == len(chunks) {
				mergedChunks[i/2] = chunks[i]
				continue
			}

			mWg.Add(1)
			go func(i int) {
				defer mWg.Done()
				a, b := chunks[i], chunks[i+1]
				start := 0
				for j := 0; j < i; j += 2 {
					start += len(chunks[j]) + len(chunks[j+1])
				}
				mergeIntoBuffer(buffer[start:], a, b, less)
				mergedChunks[i/2] = buffer[start : start+len(a)+len(b)]
			}(i)
		}
		mWg.Wait()
		chunks = mergedChunks
		toggle = !toggle
	}

	copy(data, chunks[0])
}

/* Benchmark results

BenchmarkSortStruct_Arbitrary/Arbitrary_SortStruct_1000-8	690885	108141 ns/op	76361 B/op	73 allocs/op
BenchmarkSortStruct_Arbitrary/Arbitrary_SortStruct_10000-8	89863	821504 ns/op	756302 B/op	73 allocs/op
BenchmarkSortStruct_Arbitrary/Arbitrary_SortStruct_100000-8	15216	4713591 ns/op	7228005 B/op	73 allocs/op
BenchmarkSortStruct_Arbitrary/Arbitrary_SortStruct_1000000-8	1728	41428352 ns/op	72059502 B/op	73 allocs/op
BenchmarkSortStruct_Arbitrary/Arbitrary_SortStruct_10000000-8	258	292931133 ns/op	720063072 B/op	73 allocs/op
BenchmarkSortStruct_Arbitrary/Arbitrary_SortStruct_100000000-8	26	2726562098 ns/op	7200066175 B/op	73 allocs/op
*/
