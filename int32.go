package parsort

import (
	"runtime"
	"sort"
	"sync"
)

func Int32Asc(data []int32) {
	int32Sort(data, false)
}

func Int32Desc(data []int32) {
	int32Sort(data, true)
}

func int32Sort(data []int32, reverse bool) {
	n := len(data)
	if n < 10000 {
		sort.Slice(data, func(i, j int) bool {
			return data[i] < data[j]
		})
		if reverse {
			int32Reverse(data)
		}
		return
	}

	numCPU := runtime.NumCPU()
	chunkSize := (n + numCPU - 1) / numCPU

	chunks := make([][]int32, 0, numCPU)
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
		go func(c []int32) {
			defer wg.Done()
			sort.Slice(c, func(i, j int) bool {
				return c[i] < c[j]
			})
		}(chunk)
	}
	wg.Wait()

	for len(chunks) > 1 {
		mergedCount := (len(chunks) + 1) / 2
		merged := make([][]int32, mergedCount)

		var mWg sync.WaitGroup
		for i := 0; i < len(chunks); i += 2 {
			if i+1 == len(chunks) {
				merged[i/2] = chunks[i]
				continue
			}
			mWg.Add(1)
			go func(i int) {
				defer mWg.Done()
				merged[i/2] = int32MergeSorted(chunks[i], chunks[i+1])
			}(i)
		}
		mWg.Wait()
		chunks = merged
	}

	copy(data, chunks[0])
	if reverse {
		int32Reverse(data)
	}
}

func int32MergeSorted(a, b []int32) []int32 {
	res := make([]int32, len(a)+len(b))
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

func int32Reverse(a []int32) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
