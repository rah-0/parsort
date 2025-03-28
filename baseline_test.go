package parsort

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"testing"
	"time"
)

var testSizes = []int{1000, 10000, 100000, 1000000, 10000000}

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

type person struct {
	Name string
	Age  int
}

func genPeople(n int) []person {
	a := make([]person, n)
	for i := range a {
		a[i] = person{
			Name: fmt.Sprintf("Name%d", rand.Intn(1000000)),
			Age:  rand.Intn(100),
		}
	}
	return a
}

// Method 1: sort.Slice with index-based comparator
func BenchmarkSortSlice_Arbitrary(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Arbitrary_SortSlice_"+strconv.Itoa(size), func(b *testing.B) {
			data := genPeople(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]person, len(data))
				copy(tmp, data)
				sort.Slice(tmp, func(i, j int) bool {
					return tmp[i].Age < tmp[j].Age
				})
			}
		})
	}
}

type byAge []person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// Method 2: sort.Sort with full interface
func BenchmarkSortInterface_Arbitrary(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Arbitrary_SortInterface_"+strconv.Itoa(size), func(b *testing.B) {
			data := genPeople(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]person, len(data))
				copy(tmp, data)
				sort.Sort(byAge(tmp))
			}
		})
	}
}
