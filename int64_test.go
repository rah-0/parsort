package parsort

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

func TestInt64Asc_EmptySlice(t *testing.T) {
	var data []int64
	Int64Asc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestInt64Desc_EmptySlice(t *testing.T) {
	var data []int64
	Int64Desc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestInt64Asc_SingleElement(t *testing.T) {
	data := []int64{42}
	Int64Asc(data)
	if data[0] != 42 {
		t.Errorf("expected [42], got %v", data)
	}
}

func TestInt64Desc_SingleElement(t *testing.T) {
	data := []int64{42}
	Int64Desc(data)
	if data[0] != 42 {
		t.Errorf("expected [42], got %v", data)
	}
}

func TestInt64Asc_AllEqual(t *testing.T) {
	data := make([]int64, 1000)
	for i := range data {
		data[i] = 7
	}
	Int64Asc(data)
	for _, v := range data {
		if v != 7 {
			t.Errorf("expected all elements to be 7, got %v", data)
			break
		}
	}
}

func TestInt64Desc_AllEqual(t *testing.T) {
	data := make([]int64, 1000)
	for i := range data {
		data[i] = 7
	}
	Int64Desc(data)
	for _, v := range data {
		if v != 7 {
			t.Errorf("expected all elements to be 7, got %v", data)
			break
		}
	}
}

func TestInt64Asc_AlreadySorted(t *testing.T) {
	data := make([]int64, 10000)
	for i := range data {
		data[i] = int64(i)
	}
	Int64Asc(data)
	for i := 0; i < len(data); i++ {
		if data[i] != int64(i) {
			t.Errorf("expected already sorted slice to remain unchanged")
			break
		}
	}
}

func TestInt64Desc_AlreadySorted(t *testing.T) {
	n := 10000
	data := make([]int64, n)
	for i := range data {
		data[i] = int64(n - i - 1)
	}
	Int64Desc(data)
	for i := 1; i < len(data); i++ {
		if data[i-1] < data[i] {
			t.Errorf("expected descending slice to remain unchanged")
			break
		}
	}
}

func TestInt64Asc_ReverseSorted(t *testing.T) {
	n := 10000
	data := make([]int64, n)
	for i := range data {
		data[i] = int64(n - i)
	}
	Int64Asc(data)
	for i := 1; i < n; i++ {
		if data[i-1] > data[i] {
			t.Errorf("slice not sorted in ascending order")
			break
		}
	}
}

func TestInt64Desc_ReverseSorted(t *testing.T) {
	n := 10000
	data := make([]int64, n)
	for i := range data {
		data[i] = int64(i)
	}
	Int64Desc(data)
	for i := 1; i < n; i++ {
		if data[i-1] < data[i] {
			t.Errorf("slice not sorted in descending order")
			break
		}
	}
}

func TestInt64Asc_SmallRandom(t *testing.T) {
	data := genInt64s(1000)
	expected := append([]int64(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Int64Asc(data)
	if !int64SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small ascending slice")
	}
}

func TestInt64Desc_SmallRandom(t *testing.T) {
	data := genInt64s(1000)
	expected := append([]int64(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] > expected[j] })
	Int64Desc(data)
	if !int64SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small descending slice")
	}
}

func TestInt64Asc_LargeRandom(t *testing.T) {
	data := genInt64s(2000000)
	expected := append([]int64(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Int64Asc(data)
	if !int64SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large slice")
	}
}

func TestInt64Desc_LargeRandom(t *testing.T) {
	data := genInt64s(2000000)
	expected := append([]int64(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] > expected[j] })
	Int64Desc(data)
	if !int64SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large descending slice")
	}
}

func int64SlicesEqual(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func genInt64s(n int) []int64 {
	a := make([]int64, n)
	for i := range a {
		a[i] = rand.Int63()
	}
	return a
}

func BenchmarkParsortInt64Asc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Asc_Int64_"+strconv.Itoa(size), func(b *testing.B) {
			data := genInt64s(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]int64, len(data))
				copy(tmp, data)
				Int64Asc(tmp)
			}
		})
	}
}

func BenchmarkParsortInt64Desc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Desc_Int64_"+strconv.Itoa(size), func(b *testing.B) {
			data := genInt64s(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]int64, len(data))
				copy(tmp, data)
				Int64Desc(tmp)
			}
		})
	}
}
