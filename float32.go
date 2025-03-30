package main

import (
	"sort"
	"sync"
)

func Float32Asc(data []float32) {
	float32Sort(data, false)
}

func Float32Desc(data []float32) {
	float32Sort(data, true)
}

func float32Sort(data []float32, reverse bool) {
	n := len(data)
	if n < Float32MinParallelSize {
		sort.Slice(data, func(i, j int) bool {
			return data[i] < data[j]
		})
		if reverse {
			float32Reverse(data)
		}
		return
	}

	chunkSize := (n + CoreCount - 1) / CoreCount

	chunks := make([][]float32, 0, CoreCount)
	for i := 0; i < n; i += chunkSize {
		end := i + chunkSize
		if end > n {
			end = n
		}
		chunks = append(chunks, data[i:end])
	}

	var wg sync.WaitGroup
	for _, chunk := range chunks {
		wg.Add(1)
		go func(c []float32) {
			defer wg.Done()
			sort.Slice(c, func(i, j int) bool {
				return c[i] < c[j]
			})
		}(chunk)
	}
	wg.Wait()

	for len(chunks) > 1 {
		mergedCount := (len(chunks) + 1) / 2
		merged := make([][]float32, mergedCount)

		var mWg sync.WaitGroup
		for i := 0; i < len(chunks); i += 2 {
			if i+1 == len(chunks) {
				merged[i/2] = chunks[i]
				continue
			}
			mWg.Add(1)
			go func(i int) {
				defer mWg.Done()
				merged[i/2] = float32MergeSorted(chunks[i], chunks[i+1])
			}(i)
		}
		mWg.Wait()
		chunks = merged
	}

	copy(data, chunks[0])
	if reverse {
		float32Reverse(data)
	}
}

func float32MergeSorted(a, b []float32) []float32 {
	res := make([]float32, len(a)+len(b))
	i, j, k := 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			res[k] = a[i]
			i++
		} else {
			res[k] = b[j]
			j++
		}
		k++
	}
	for i < len(a) {
		res[k] = a[i]
		i++
		k++
	}
	for j < len(b) {
		res[k] = b[j]
		j++
		k++
	}
	return res
}

func float32Reverse(a []float32) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
