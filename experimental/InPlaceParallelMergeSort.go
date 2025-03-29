package experimental

import (
	"sort"
	"sync"
)

type chunk struct{ start, end int }

func InPlaceParallelMergeSort[T any](data []T, less func(a, b T) bool) {
	n := len(data)
	if n <= 1 {
		return
	}

	chunkSize := (n + CoreCount - 1) / CoreCount
	chunks := make([]chunk, 0, CoreCount)

	for i := 0; i < n; i += chunkSize {
		end := i + chunkSize
		if end > n {
			end = n
		}
		chunks = append(chunks, chunk{i, end})
	}

	var wg sync.WaitGroup
	prefetch := make(chan struct{}, CoreCount)
	for _, ch := range chunks {
		wg.Add(1)
		prefetch <- struct{}{}
		go func(start, end int) {
			defer wg.Done()
			defer func() { <-prefetch }()
			sort.Slice(data[start:end], func(i, j int) bool {
				return less(data[start+i], data[start+j])
			})
		}(ch.start, ch.end)
	}
	wg.Wait()

	dst := make([]T, n)
	src := data

	for len(chunks) > 1 {
		merged := make([]chunk, 0, (len(chunks)+1)/2)
		var mWg sync.WaitGroup
		mergePrefetch := make(chan struct{}, CoreCount)

		for i := 0; i < len(chunks); i += 2 {
			if i+1 == len(chunks) {
				copy(dst[chunks[i].start:chunks[i].end], src[chunks[i].start:chunks[i].end])
				merged = append(merged, chunks[i])
				continue
			}

			a, b := chunks[i], chunks[i+1]
			outStart := a.start
			outEnd := b.end
			merged = append(merged, chunk{outStart, outEnd})

			mWg.Add(1)
			mergePrefetch <- struct{}{}
			go func(a, b chunk) {
				defer mWg.Done()
				defer func() { <-mergePrefetch }()
				i, j, k := a.start, b.start, a.start
				for i < a.end && j < b.end {
					if less(src[i], src[j]) {
						dst[k] = src[i]
						i++
					} else {
						dst[k] = src[j]
						j++
					}
					k++
				}
				copy(dst[k:], src[i:a.end])
				copy(dst[k+(a.end-i):], src[j:b.end])
			}(a, b)
		}
		mWg.Wait()
		src, dst = dst, src
		chunks = merged
	}

	if &src[0] != &data[0] {
		copy(data, src)
	}
}

/* Benchmark results

BenchmarkSortStruct_Arbitrary/Arbitrary_SortStruct_1000-8	641217	112847 ns/op	52320 B/op	72 allocs/op
BenchmarkSortStruct_Arbitrary/Arbitrary_SortStruct_10000-8	101338	711605 ns/op	494691 B/op	72 allocs/op
BenchmarkSortStruct_Arbitrary/Arbitrary_SortStruct_100000-8	18980	3834272 ns/op	4803692 B/op	72 allocs/op
BenchmarkSortStruct_Arbitrary/Arbitrary_SortStruct_1000000-8	2244	31901014 ns/op	48008297 B/op	72 allocs/op
BenchmarkSortStruct_Arbitrary/Arbitrary_SortStruct_10000000-8	279	250564309 ns/op	480005222 B/op	72 allocs/op
BenchmarkSortStruct_Arbitrary/Arbitrary_SortStruct_100000000-8	28	2580231136 ns/op	4800007267 B/op	72 allocs/op
*/
