package parsort

import (
	"sort"
	"sync"
)

type chunk struct{ start, end int }

// structSort sorts a slice using parallel stable sorting and in-place merging.
func structSort[T any](data []T, less func(a, b T) bool) {
	n := len(data)
	if n <= 1 {
		return
	}

	if n < 10000 {
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
						// Equal elements, preserve stability by choosing left first
						dst[k] = src[i]
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

// StructAsc sorts a slice of structs in ascending order based on the provided less function.
func StructAsc[T any](data []T, less func(a, b T) bool) {
	structSort(data, less)
}

// StructDesc sorts a slice of structs in descending order based on the provided less function.
func StructDesc[T any](data []T, less func(a, b T) bool) {
	structSort(data, func(a, b T) bool {
		return less(b, a)
	})
}
