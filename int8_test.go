package parsort

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

func TestInt8Asc_EmptySlice(t *testing.T) {
	var data []int8
	Int8Asc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestInt8Desc_EmptySlice(t *testing.T) {
	var data []int8
	Int8Desc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestInt8Asc_SingleElement(t *testing.T) {
	data := []int8{42}
	Int8Asc(data)
	if data[0] != 42 {
		t.Errorf("expected [42], got %v", data)
	}
}

func TestInt8Desc_SingleElement(t *testing.T) {
	data := []int8{42}
	Int8Desc(data)
	if data[0] != 42 {
		t.Errorf("expected [42], got %v", data)
	}
}

func TestInt8Asc_AllEqual(t *testing.T) {
	data := make([]int8, 1000)
	for i := range data {
		data[i] = 7
	}
	Int8Asc(data)
	for _, v := range data {
		if v != 7 {
			t.Errorf("expected all elements to be 7, got %v", data)
			break
		}
	}
}

func TestInt8Desc_AllEqual(t *testing.T) {
	data := make([]int8, 1000)
	for i := range data {
		data[i] = 7
	}
	Int8Desc(data)
	for _, v := range data {
		if v != 7 {
			t.Errorf("expected all elements to be 7, got %v", data)
			break
		}
	}
}

func TestInt8Asc_AlreadySorted(t *testing.T) {
	data := make([]int8, 127)
	for i := range data {
		data[i] = int8(i)
	}
	Int8Asc(data)
	for i := 0; i < len(data); i++ {
		if data[i] != int8(i) {
			t.Errorf("expected already sorted slice to remain unchanged")
			break
		}
	}
}

func TestInt8Desc_AlreadySorted(t *testing.T) {
	n := 127
	data := make([]int8, n)
	for i := range data {
		data[i] = int8(n - i - 1)
	}
	Int8Desc(data)
	for i := 1; i < len(data); i++ {
		if data[i-1] < data[i] {
			t.Errorf("expected already descending slice to remain unchanged")
			break
		}
	}
}

func TestInt8Asc_ReverseSorted(t *testing.T) {
	n := 127
	data := make([]int8, n)
	for i := range data {
		data[i] = int8(n - i)
	}
	Int8Asc(data)
	for i := 1; i < n; i++ {
		if data[i-1] > data[i] {
			t.Errorf("slice not sorted in ascending order")
			break
		}
	}
}

func TestInt8Desc_ReverseSorted(t *testing.T) {
	n := 127
	data := make([]int8, n)
	for i := range data {
		data[i] = int8(i)
	}
	Int8Desc(data)
	for i := 1; i < n; i++ {
		if data[i-1] < data[i] {
			t.Errorf("slice not sorted in descending order")
			break
		}
	}
}

func TestInt8Asc_LargeRandom(t *testing.T) {
	data := genInt8s(2000000)
	expected := append([]int8(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Int8Asc(data)
	if !int8SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large slice")
	}
}

func TestInt8Desc_LargeRandom(t *testing.T) {
	data := genInt8s(2000000)
	expected := append([]int8(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] > expected[j] })
	Int8Desc(data)
	if !int8SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large descending slice")
	}
}

func TestInt8Asc_SmallRandom(t *testing.T) {
	data := genInt8s(1000)
	expected := append([]int8(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Int8Asc(data)
	if !int8SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small ascending slice")
	}
}

func TestInt8Desc_SmallRandom(t *testing.T) {
	data := genInt8s(1000)
	expected := append([]int8(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] > expected[j] })
	Int8Desc(data)
	if !int8SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small descending slice")
	}
}

func int8SlicesEqual(a, b []int8) bool {
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

func genInt8s(n int) []int8 {
	a := make([]int8, n)
	for i := range a {
		a[i] = int8(rand.Intn(256) - 128)
	}
	return a
}

func BenchmarkParsortInt8Asc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Asc_Int8_"+strconv.Itoa(size), func(b *testing.B) {
			data := genInt8s(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]int8, len(data))
				copy(tmp, data)
				Int8Asc(tmp)
			}
		})
	}
}

func BenchmarkParsortInt8Desc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Desc_Int8_"+strconv.Itoa(size), func(b *testing.B) {
			data := genInt8s(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]int8, len(data))
				copy(tmp, data)
				Int8Desc(tmp)
			}
		})
	}
}
