package main

import (
	"sort"
	"strconv"
	"testing"
)

func TestUint16Asc_EmptySlice(t *testing.T) {
	var data []uint16
	Uint16Asc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestUint16Desc_EmptySlice(t *testing.T) {
	var data []uint16
	Uint16Desc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestUint16Asc_SingleElement(t *testing.T) {
	data := []uint16{42}
	Uint16Asc(data)
	if data[0] != 42 {
		t.Errorf("expected [42], got %v", data)
	}
}

func TestUint16Desc_SingleElement(t *testing.T) {
	data := []uint16{42}
	Uint16Desc(data)
	if data[0] != 42 {
		t.Errorf("expected [42], got %v", data)
	}
}

func TestUint16Asc_AllEqual(t *testing.T) {
	data := make([]uint16, 1000)
	for i := range data {
		data[i] = 123
	}
	Uint16Asc(data)
	for _, v := range data {
		if v != 123 {
			t.Errorf("expected all elements to be 123, got %v", data)
			break
		}
	}
}

func TestUint16Desc_AllEqual(t *testing.T) {
	data := make([]uint16, 1000)
	for i := range data {
		data[i] = 123
	}
	Uint16Desc(data)
	for _, v := range data {
		if v != 123 {
			t.Errorf("expected all elements to be 123, got %v", data)
			break
		}
	}
}

func TestUint16Asc_AlreadySorted(t *testing.T) {
	data := make([]uint16, 10000)
	for i := range data {
		data[i] = uint16(i)
	}
	Uint16Asc(data)
	for i := 0; i < len(data); i++ {
		if data[i] != uint16(i) {
			t.Errorf("expected already sorted slice to remain unchanged")
			break
		}
	}
}

func TestUint16Desc_AlreadySorted(t *testing.T) {
	n := 10000
	data := make([]uint16, n)
	for i := range data {
		data[i] = uint16(n - i - 1)
	}
	Uint16Desc(data)
	for i := 1; i < len(data); i++ {
		if data[i-1] < data[i] {
			t.Errorf("expected descending slice to remain unchanged")
			break
		}
	}
}

func TestUint16Asc_ReverseSorted(t *testing.T) {
	n := 10000
	data := make([]uint16, n)
	for i := range data {
		data[i] = uint16(n - i)
	}
	Uint16Asc(data)
	for i := 1; i < n; i++ {
		if data[i-1] > data[i] {
			t.Errorf("slice not sorted in ascending order")
			break
		}
	}
}

func TestUint16Desc_ReverseSorted(t *testing.T) {
	n := 10000
	data := make([]uint16, n)
	for i := range data {
		data[i] = uint16(i)
	}
	Uint16Desc(data)
	for i := 1; i < n; i++ {
		if data[i-1] < data[i] {
			t.Errorf("slice not sorted in descending order")
			break
		}
	}
}

func TestUint16Asc_SmallRandom(t *testing.T) {
	data := genUint16s(1000)
	expected := append([]uint16(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Uint16Asc(data)
	if !uint16SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small ascending slice")
	}
}

func TestUint16Desc_SmallRandom(t *testing.T) {
	data := genUint16s(1000)
	expected := append([]uint16(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] > expected[j] })
	Uint16Desc(data)
	if !uint16SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small descending slice")
	}
}

func TestUint16Asc_LargeRandom(t *testing.T) {
	data := genUint16s(2000000)
	expected := append([]uint16(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Uint16Asc(data)
	if !uint16SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large slice")
	}
}

func TestUint16Desc_LargeRandom(t *testing.T) {
	data := genUint16s(2000000)
	expected := append([]uint16(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] > expected[j] })
	Uint16Desc(data)
	if !uint16SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large descending slice")
	}
}

func BenchmarkParsortUint16Asc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Asc_Uint16_"+strconv.Itoa(size), func(b *testing.B) {
			data := genUint16s(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]uint16, len(data))
				copy(tmp, data)
				Uint16Asc(tmp)
			}
		})
	}
}

func BenchmarkParsortUint16Desc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Desc_Uint16_"+strconv.Itoa(size), func(b *testing.B) {
			data := genUint16s(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]uint16, len(data))
				copy(tmp, data)
				Uint16Desc(tmp)
			}
		})
	}
}
