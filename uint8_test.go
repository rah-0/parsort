package main

import (
	"sort"
	"strconv"
	"testing"
)

func TestUint8Asc_EmptySlice(t *testing.T) {
	var data []uint8
	Uint8Asc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestUint8Desc_EmptySlice(t *testing.T) {
	var data []uint8
	Uint8Desc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestUint8Asc_SingleElement(t *testing.T) {
	data := []uint8{42}
	Uint8Asc(data)
	if data[0] != 42 {
		t.Errorf("expected [42], got %v", data)
	}
}

func TestUint8Desc_SingleElement(t *testing.T) {
	data := []uint8{42}
	Uint8Desc(data)
	if data[0] != 42 {
		t.Errorf("expected [42], got %v", data)
	}
}

func TestUint8Asc_AllEqual(t *testing.T) {
	data := make([]uint8, 1000)
	for i := range data {
		data[i] = 7
	}
	Uint8Asc(data)
	for _, v := range data {
		if v != 7 {
			t.Errorf("expected all elements to be 7, got %v", data)
			break
		}
	}
}

func TestUint8Desc_AllEqual(t *testing.T) {
	data := make([]uint8, 1000)
	for i := range data {
		data[i] = 7
	}
	Uint8Desc(data)
	for _, v := range data {
		if v != 7 {
			t.Errorf("expected all elements to be 7, got %v", data)
			break
		}
	}
}

func TestUint8Asc_AlreadySorted(t *testing.T) {
	data := make([]uint8, 255)
	for i := range data {
		data[i] = uint8(i)
	}
	Uint8Asc(data)
	for i := 0; i < len(data); i++ {
		if data[i] != uint8(i) {
			t.Errorf("expected already sorted slice to remain unchanged")
			break
		}
	}
}

func TestUint8Desc_AlreadySorted(t *testing.T) {
	n := 255
	data := make([]uint8, n)
	for i := range data {
		data[i] = uint8(n - i - 1)
	}
	Uint8Desc(data)
	for i := 1; i < len(data); i++ {
		if data[i-1] < data[i] {
			t.Errorf("expected descending slice to remain unchanged")
			break
		}
	}
}

func TestUint8Asc_ReverseSorted(t *testing.T) {
	n := 255
	data := make([]uint8, n)
	for i := range data {
		data[i] = uint8(n - i)
	}
	Uint8Asc(data)
	for i := 1; i < n; i++ {
		if data[i-1] > data[i] {
			t.Errorf("slice not sorted in ascending order")
			break
		}
	}
}

func TestUint8Desc_ReverseSorted(t *testing.T) {
	n := 255
	data := make([]uint8, n)
	for i := range data {
		data[i] = uint8(i)
	}
	Uint8Desc(data)
	for i := 1; i < n; i++ {
		if data[i-1] < data[i] {
			t.Errorf("slice not sorted in descending order")
			break
		}
	}
}

func TestUint8Asc_SmallRandom(t *testing.T) {
	data := genUint8s(1000)
	expected := append([]uint8(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Uint8Asc(data)
	if !uint8SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small ascending slice")
	}
}

func TestUint8Desc_SmallRandom(t *testing.T) {
	data := genUint8s(1000)
	expected := append([]uint8(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] > expected[j] })
	Uint8Desc(data)
	if !uint8SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small descending slice")
	}
}

func TestUint8Asc_LargeRandom(t *testing.T) {
	data := genUint8s(2000000)
	expected := append([]uint8(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Uint8Asc(data)
	if !uint8SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large slice")
	}
}

func TestUint8Desc_LargeRandom(t *testing.T) {
	data := genUint8s(2000000)
	expected := append([]uint8(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] > expected[j] })
	Uint8Desc(data)
	if !uint8SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large descending slice")
	}
}

func BenchmarkParsortUint8Asc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Asc_Uint8_"+strconv.Itoa(size), func(b *testing.B) {
			data := genUint8s(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]uint8, len(data))
				copy(tmp, data)
				Uint8Asc(tmp)
			}
		})
	}
}

func BenchmarkParsortUint8Desc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Desc_Uint8_"+strconv.Itoa(size), func(b *testing.B) {
			data := genUint8s(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]uint8, len(data))
				copy(tmp, data)
				Uint8Desc(tmp)
			}
		})
	}
}
