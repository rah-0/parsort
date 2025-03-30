package main

import (
	"sort"
	"strconv"
	"testing"
)

func TestUint32Asc_EmptySlice(t *testing.T) {
	var data []uint32
	Uint32Asc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestUint32Desc_EmptySlice(t *testing.T) {
	var data []uint32
	Uint32Desc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestUint32Asc_SingleElement(t *testing.T) {
	data := []uint32{42}
	Uint32Asc(data)
	if data[0] != 42 {
		t.Errorf("expected [42], got %v", data)
	}
}

func TestUint32Desc_SingleElement(t *testing.T) {
	data := []uint32{42}
	Uint32Desc(data)
	if data[0] != 42 {
		t.Errorf("expected [42], got %v", data)
	}
}

func TestUint32Asc_AllEqual(t *testing.T) {
	data := make([]uint32, 1000)
	for i := range data {
		data[i] = 7
	}
	Uint32Asc(data)
	for _, v := range data {
		if v != 7 {
			t.Errorf("expected all elements to be 7, got %v", data)
			break
		}
	}
}

func TestUint32Desc_AllEqual(t *testing.T) {
	data := make([]uint32, 1000)
	for i := range data {
		data[i] = 7
	}
	Uint32Desc(data)
	for _, v := range data {
		if v != 7 {
			t.Errorf("expected all elements to be 7, got %v", data)
			break
		}
	}
}

func TestUint32Asc_AlreadySorted(t *testing.T) {
	data := make([]uint32, 10000)
	for i := range data {
		data[i] = uint32(i)
	}
	Uint32Asc(data)
	for i := 0; i < len(data); i++ {
		if data[i] != uint32(i) {
			t.Errorf("expected already sorted slice to remain unchanged")
			break
		}
	}
}

func TestUint32Desc_AlreadySorted(t *testing.T) {
	n := 10000
	data := make([]uint32, n)
	for i := range data {
		data[i] = uint32(n - i - 1)
	}
	Uint32Desc(data)
	for i := 1; i < len(data); i++ {
		if data[i-1] < data[i] {
			t.Errorf("expected descending slice to remain unchanged")
			break
		}
	}
}

func TestUint32Asc_ReverseSorted(t *testing.T) {
	n := 10000
	data := make([]uint32, n)
	for i := range data {
		data[i] = uint32(n - i)
	}
	Uint32Asc(data)
	for i := 1; i < n; i++ {
		if data[i-1] > data[i] {
			t.Errorf("slice not sorted in ascending order")
			break
		}
	}
}

func TestUint32Desc_ReverseSorted(t *testing.T) {
	n := 10000
	data := make([]uint32, n)
	for i := range data {
		data[i] = uint32(i)
	}
	Uint32Desc(data)
	for i := 1; i < n; i++ {
		if data[i-1] < data[i] {
			t.Errorf("slice not sorted in descending order")
			break
		}
	}
}

func TestUint32Asc_SmallRandom(t *testing.T) {
	data := genUint32s(1000)
	expected := append([]uint32(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Uint32Asc(data)
	if !uint32SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small ascending slice")
	}
}

func TestUint32Desc_SmallRandom(t *testing.T) {
	data := genUint32s(1000)
	expected := append([]uint32(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] > expected[j] })
	Uint32Desc(data)
	if !uint32SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small descending slice")
	}
}

func TestUint32Asc_LargeRandom(t *testing.T) {
	data := genUint32s(2000000)
	expected := append([]uint32(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Uint32Asc(data)
	if !uint32SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large slice")
	}
}

func TestUint32Desc_LargeRandom(t *testing.T) {
	data := genUint32s(2000000)
	expected := append([]uint32(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] > expected[j] })
	Uint32Desc(data)
	if !uint32SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large descending slice")
	}
}

func BenchmarkParsortUint32Asc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Asc_Uint32_"+strconv.Itoa(size), func(b *testing.B) {
			data := genUint32s(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]uint32, len(data))
				copy(tmp, data)
				Uint32Asc(tmp)
			}
		})
	}
}

func BenchmarkParsortUint32Desc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Desc_Uint32_"+strconv.Itoa(size), func(b *testing.B) {
			data := genUint32s(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]uint32, len(data))
				copy(tmp, data)
				Uint32Desc(tmp)
			}
		})
	}
}
