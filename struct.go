package main

import (
	"sort"
	"sync"
)

type chunk struct{ start, end int }

// structSortUnstable sorts a slice using parallel unstable sorting and in-place merging.
func structSortUnstable[T any](data []T, less func(a, b T) bool) {
	n := len(data)
	if n <= 1 {
		return
	}

	if n < StructMinParallelSize {
		sort.Slice(data, func(i, j int) bool {
			return less(data[i], data[j])
		})
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
	for _, ch := range chunks {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
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
			go func(a, b chunk) {
				defer mWg.Done()
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

// structSortStable sorts a slice using parallel stable sorting and in-place merging.
func structSortStable[T any](data []T, less func(a, b T) bool) {
	n := len(data)
	if n <= 1 {
		return
	}

	if n < StructMinParallelSize {
		sort.SliceStable(data, func(i, j int) bool {
			return less(data[i], data[j])
		})
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
	for _, ch := range chunks {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			sort.SliceStable(data[start:end], func(i, j int) bool {
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
			go func(a, b chunk) {
				defer mWg.Done()
				i, j, k := a.start, b.start, a.start
				for i < a.end && j < b.end {
					switch {
					case less(src[i], src[j]):
						dst[k] = src[i]
						i++
					case less(src[j], src[i]):
						dst[k] = src[j]
						j++
					default:
						dst[k] = src[i] // preserve left-side element
						i++
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

// StructAsc sorts a slice of structs in ascending order using unstable sort.
func StructAsc[T any](data []T, less func(a, b T) bool) {
	structSortUnstable(data, less)
}

// StructDesc sorts a slice of structs in descending order using unstable sort.
func StructDesc[T any](data []T, less func(a, b T) bool) {
	structSortUnstable(data, func(a, b T) bool {
		return less(b, a)
	})
}

// StructAscStable sorts a slice of structs in ascending order using stable sort.
func StructAscStable[T any](data []T, less func(a, b T) bool) {
	structSortStable(data, less)
}

// StructDescStable sorts a slice of structs in descending order using stable sort.
func StructDescStable[T any](data []T, less func(a, b T) bool) {
	structSortStable(data, func(a, b T) bool {
		return less(b, a)
	})
}
