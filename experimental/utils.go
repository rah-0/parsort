package experimental

import (
	"runtime"
)

var (
	CoreCount = runtime.NumCPU()
)

func mergeIntoBuffer[T any](dst, a, b []T, less func(a, b T) bool) {
	i, j, k := 0, 0, 0
	for i < len(a) && j < len(b) {
		if less(a[i], b[j]) {
			dst[k] = a[i]
			i++
		} else {
			dst[k] = b[j]
			j++
		}
		k++
	}
	for i < len(a) {
		dst[k] = a[i]
		i++
		k++
	}
	for j < len(b) {
		dst[k] = b[j]
		j++
		k++
	}
}
