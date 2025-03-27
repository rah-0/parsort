package parsort

import (
	"sort"
	"strconv"
	"testing"
	"time"
)

var testSizes = []int{1_000, 10_000, 100_000, 1_000_000, 10_000_000}

func BenchmarkBaselineSortIntsAsc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Asc_Int_"+strconv.Itoa(size), func(b *testing.B) {
			data := genInts(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]int, len(data))
				copy(tmp, data)
				sort.Ints(tmp)
			}
		})
	}
}

func BenchmarkBaselineSortIntsDesc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Desc_Int_"+strconv.Itoa(size), func(b *testing.B) {
			data := genInts(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]int, len(data))
				copy(tmp, data)
				sort.Sort(sort.Reverse(sort.IntSlice(tmp)))
			}
		})
	}
}

func BenchmarkBaselineSortFloatsAsc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Asc_Float_"+strconv.Itoa(size), func(b *testing.B) {
			data := genFloats(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]float64, len(data))
				copy(tmp, data)
				sort.Float64s(tmp)
			}
		})
	}
}

func BenchmarkBaselineSortFloatsDesc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Desc_Float_"+strconv.Itoa(size), func(b *testing.B) {
			data := genFloats(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]float64, len(data))
				copy(tmp, data)
				sort.Sort(sort.Reverse(sort.Float64Slice(tmp)))
			}
		})
	}
}

func BenchmarkBaselineSortTimes(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Sort_Time_"+strconv.Itoa(size), func(b *testing.B) {
			data := genTimes(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]time.Time, len(data))
				copy(tmp, data)
				sort.Slice(tmp, func(i, j int) bool {
					return tmp[i].Before(tmp[j])
				})
			}
		})
	}
}

func BenchmarkBaselineSortStringsAsc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Asc_String_"+strconv.Itoa(size), func(b *testing.B) {
			data := genStrings(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]string, len(data))
				copy(tmp, data)
				sort.Strings(tmp)
			}
		})
	}
}

func BenchmarkBaselineSortStringsDesc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Desc_String_"+strconv.Itoa(size), func(b *testing.B) {
			data := genStrings(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]string, len(data))
				copy(tmp, data)
				sort.Sort(sort.Reverse(sort.StringSlice(tmp)))
			}
		})
	}
}
