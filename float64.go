package parsort

import (
	"sort"
	"sync"
)

func Float64Asc(data []float64) {
	float64Sort(data, false)
}

func Float64Desc(data []float64) {
	float64Sort(data, true)
}

func float64Sort(data []float64, reverse bool) {
	n := len(data)
	if n < Float64MinParallelSize {
		sort.Float64s(data)
		if reverse {
			float64Reverse(data)
		}
		return
	}

	chunkSize := (n + CoreCount - 1) / CoreCount

	chunks := make([][]float64, 0, CoreCount)
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
		go func(c []float64) {
			defer wg.Done()
			sort.Float64s(c)
		}(chunk)
	}
	wg.Wait()

	// Parallel merging loop
	for len(chunks) > 1 {
		mergedCount := (len(chunks) + 1) / 2
		merged := make([][]float64, mergedCount)

		var mWg sync.WaitGroup
		for i := 0; i < len(chunks); i += 2 {
			if i+1 == len(chunks) {
				merged[i/2] = chunks[i]
				continue
			}
			mWg.Add(1)
			go func(i int) {
				defer mWg.Done()
				merged[i/2] = float64MergeSorted(chunks[i], chunks[i+1])
			}(i)
		}
		mWg.Wait()
		chunks = merged
	}

	copy(data, chunks[0])
	if reverse {
		float64Reverse(data)
	}
}

func float64MergeSorted(a, b []float64) []float64 {
	res := make([]float64, len(a)+len(b))
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

func float64Reverse(a []float64) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
