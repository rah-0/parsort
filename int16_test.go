package parsort

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

func TestInt16Asc_EmptySlice(t *testing.T) {
	var data []int16
	Int16Asc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestInt16Desc_EmptySlice(t *testing.T) {
	var data []int16
	Int16Desc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestInt16Asc_SingleElement(t *testing.T) {
	data := []int16{42}
	Int16Asc(data)
	if data[0] != 42 {
		t.Errorf("expected [42], got %v", data)
	}
}

func TestInt16Desc_SingleElement(t *testing.T) {
	data := []int16{42}
	Int16Desc(data)
	if data[0] != 42 {
		t.Errorf("expected [42], got %v", data)
	}
}

func TestInt16Asc_AllEqual(t *testing.T) {
	data := make([]int16, 1000)
	for i := range data {
		data[i] = 7
	}
	Int16Asc(data)
	for _, v := range data {
		if v != 7 {
			t.Errorf("expected all elements to be 7, got %v", data)
			break
		}
	}
}

func TestInt16Desc_AllEqual(t *testing.T) {
	data := make([]int16, 1000)
	for i := range data {
		data[i] = 7
	}
	Int16Desc(data)
	for _, v := range data {
		if v != 7 {
			t.Errorf("expected all elements to be 7, got %v", data)
			break
		}
	}
}

func TestInt16Asc_AlreadySorted(t *testing.T) {
	data := make([]int16, 10000)
	for i := range data {
		data[i] = int16(i)
	}
	Int16Asc(data)
	for i := 0; i < len(data); i++ {
		if data[i] != int16(i) {
			t.Errorf("expected already sorted slice to remain unchanged")
			break
		}
	}
}

func TestInt16Desc_AlreadySorted(t *testing.T) {
	n := 10000
	data := make([]int16, n)
	for i := range data {
		data[i] = int16(n - i - 1)
	}
	Int16Desc(data)
	for i := 1; i < len(data); i++ {
		if data[i-1] < data[i] {
			t.Errorf("expected descending slice to remain unchanged")
			break
		}
	}
}

func TestInt16Asc_ReverseSorted(t *testing.T) {
	n := 10000
	data := make([]int16, n)
	for i := range data {
		data[i] = int16(n - i)
	}
	Int16Asc(data)
	for i := 1; i < n; i++ {
		if data[i-1] > data[i] {
			t.Errorf("slice not sorted in ascending order")
			break
		}
	}
}

func TestInt16Desc_ReverseSorted(t *testing.T) {
	n := 10000
	data := make([]int16, n)
	for i := range data {
		data[i] = int16(i)
	}
	Int16Desc(data)
	for i := 1; i < n; i++ {
		if data[i-1] < data[i] {
			t.Errorf("slice not sorted in descending order")
			break
		}
	}
}

func TestInt16Asc_SmallRandom(t *testing.T) {
	data := genInt16s(1000)
	expected := append([]int16(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Int16Asc(data)
	if !int16SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small ascending slice")
	}
}

func TestInt16Desc_SmallRandom(t *testing.T) {
	data := genInt16s(1000)
	expected := append([]int16(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] > expected[j] })
	Int16Desc(data)
	if !int16SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small descending slice")
	}
}

func TestInt16Asc_LargeRandom(t *testing.T) {
	data := genInt16s(2_000_000)
	expected := append([]int16(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Int16Asc(data)
	if !int16SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large slice")
	}
}

func TestInt16Desc_LargeRandom(t *testing.T) {
	data := genInt16s(2_000_000)
	expected := append([]int16(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] > expected[j] })
	Int16Desc(data)
	if !int16SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large descending slice")
	}
}

func int16SlicesEqual(a, b []int16) bool {
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

func genInt16s(n int) []int16 {
	a := make([]int16, n)
	for i := range a {
		a[i] = int16(rand.Intn(65536) - 32768)
	}
	return a
}

func BenchmarkParsortInt16Asc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Asc_Int16_"+strconv.Itoa(size), func(b *testing.B) {
			data := genInt16s(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]int16, len(data))
				copy(tmp, data)
				Int16Asc(tmp)
			}
		})
	}
}

func BenchmarkParsortInt16Desc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Desc_Int16_"+strconv.Itoa(size), func(b *testing.B) {
			data := genInt16s(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]int16, len(data))
				copy(tmp, data)
				Int16Desc(tmp)
			}
		})
	}
}
