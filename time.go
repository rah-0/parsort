package main

import (
	"sort"
	"sync"
	"time"
)

func TimeAsc(data []time.Time) {
	timeSort(data, false)
}

func TimeDesc(data []time.Time) {
	timeSort(data, true)
}

func timeSort(data []time.Time, reverse bool) {
	n := len(data)
	if n < TimeMinParallelSize {
		sort.Slice(data, func(i, j int) bool {
			return data[i].Before(data[j])
		})
		if reverse {
			timeReverse(data)
		}
		return
	}

	chunkSize := (n + CoreCount - 1) / CoreCount

	chunks := make([][]time.Time, 0, CoreCount)
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
		go func(c []time.Time) {
			defer wg.Done()
			sort.Slice(c, func(i, j int) bool {
				return c[i].Before(c[j])
			})
		}(chunk)
	}
	wg.Wait()

	for len(chunks) > 1 {
		mergedCount := (len(chunks) + 1) / 2
		merged := make([][]time.Time, mergedCount)

		var mWg sync.WaitGroup
		for i := 0; i < len(chunks); i += 2 {
			if i+1 == len(chunks) {
				merged[i/2] = chunks[i]
				continue
			}
			mWg.Add(1)
			go func(i int) {
				defer mWg.Done()
				merged[i/2] = timeMergeSorted(chunks[i], chunks[i+1])
			}(i)
		}
		mWg.Wait()
		chunks = merged
	}

	copy(data, chunks[0])
	if reverse {
		timeReverse(data)
	}
}

func timeMergeSorted(a, b []time.Time) []time.Time {
	res := make([]time.Time, len(a)+len(b))
	i, j, k := 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i].Before(b[j]) || a[i].Equal(b[j]) {
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

func timeReverse(a []time.Time) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
