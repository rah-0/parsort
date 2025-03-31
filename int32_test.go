package parsort

import (
	"sort"
	"strconv"
	"testing"
)

func TestInt32Asc_EmptySlice(t *testing.T) {
	var data []int32
	Int32Asc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestInt32Desc_EmptySlice(t *testing.T) {
	var data []int32
	Int32Desc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestInt32Asc_SingleElement(t *testing.T) {
	data := []int32{42}
	Int32Asc(data)
	if data[0] != 42 {
		t.Errorf("expected [42], got %v", data)
	}
}

func TestInt32Desc_SingleElement(t *testing.T) {
	data := []int32{42}
	Int32Desc(data)
	if data[0] != 42 {
		t.Errorf("expected [42], got %v", data)
	}
}

func TestInt32Asc_AllEqual(t *testing.T) {
	data := make([]int32, 1000)
	for i := range data {
		data[i] = 7
	}
	Int32Asc(data)
	for _, v := range data {
		if v != 7 {
			t.Errorf("expected all elements to be 7, got %v", data)
			break
		}
	}
}

func TestInt32Desc_AllEqual(t *testing.T) {
	data := make([]int32, 1000)
	for i := range data {
		data[i] = 7
	}
	Int32Desc(data)
	for _, v := range data {
		if v != 7 {
			t.Errorf("expected all elements to be 7, got %v", data)
			break
		}
	}
}

func TestInt32Asc_AlreadySorted(t *testing.T) {
	data := make([]int32, 10000)
	for i := range data {
		data[i] = int32(i)
	}
	Int32Asc(data)
	for i := 0; i < len(data); i++ {
		if data[i] != int32(i) {
			t.Errorf("expected already sorted slice to remain unchanged")
			break
		}
	}
}

func TestInt32Desc_AlreadySorted(t *testing.T) {
	n := 10000
	data := make([]int32, n)
	for i := range data {
		data[i] = int32(n - i - 1)
	}
	Int32Desc(data)
	for i := 1; i < len(data); i++ {
		if data[i-1] < data[i] {
			t.Errorf("expected already descending slice to remain unchanged")
			break
		}
	}
}

func TestInt32Asc_ReverseSorted(t *testing.T) {
	n := 10000
	data := make([]int32, n)
	for i := range data {
		data[i] = int32(n - i)
	}
	Int32Asc(data)
	for i := 1; i < n; i++ {
		if data[i-1] > data[i] {
			t.Errorf("slice not sorted in ascending order")
			break
		}
	}
}

func TestInt32Desc_ReverseSorted(t *testing.T) {
	n := 10000
	data := make([]int32, n)
	for i := range data {
		data[i] = int32(i)
	}
	Int32Desc(data)
	for i := 1; i < n; i++ {
		if data[i-1] < data[i] {
			t.Errorf("slice not sorted in descending order")
			break
		}
	}
}

func TestInt32Asc_LargeRandom(t *testing.T) {
	data := genInt32s(2000000)
	expected := append([]int32(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Int32Asc(data)
	if !int32SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large slice")
	}
}

func TestInt32Desc_LargeRandom(t *testing.T) {
	data := genInt32s(2000000)
	expected := append([]int32(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] > expected[j] })
	Int32Desc(data)
	if !int32SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large descending slice")
	}
}

func TestInt32Asc_SmallRandom(t *testing.T) {
	data := genInt32s(1000)
	expected := append([]int32(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Int32Asc(data)
	if !int32SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small ascending slice")
	}
}

func TestInt32Desc_SmallRandom(t *testing.T) {
	data := genInt32s(1000)
	expected := append([]int32(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] > expected[j] })
	Int32Desc(data)
	if !int32SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small descending slice")
	}
}

func BenchmarkParsortInt32Asc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Asc_Int32_"+strconv.Itoa(size), func(b *testing.B) {
			data := genInt32s(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]int32, len(data))
				copy(tmp, data)
				Int32Asc(tmp)
			}
		})
	}
}

func BenchmarkParsortInt32Desc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Desc_Int32_"+strconv.Itoa(size), func(b *testing.B) {
			data := genInt32s(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]int32, len(data))
				copy(tmp, data)
				Int32Desc(tmp)
			}
		})
	}
}
