package main

import (
	"sort"
	"strconv"
	"testing"
)

func TestUintAsc_EmptySlice(t *testing.T) {
	var data []uint
	UintAsc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestUintDesc_EmptySlice(t *testing.T) {
	var data []uint
	UintDesc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestUintAsc_SingleElement(t *testing.T) {
	data := []uint{42}
	UintAsc(data)
	if data[0] != 42 {
		t.Errorf("expected [42], got %v", data)
	}
}

func TestUintDesc_SingleElement(t *testing.T) {
	data := []uint{42}
	UintDesc(data)
	if data[0] != 42 {
		t.Errorf("expected [42], got %v", data)
	}
}

func TestUintAsc_AllEqual(t *testing.T) {
	data := make([]uint, 1000)
	for i := range data {
		data[i] = 7
	}
	UintAsc(data)
	for _, v := range data {
		if v != 7 {
			t.Errorf("expected all elements to be 7, got %v", data)
			break
		}
	}
}

func TestUintDesc_AllEqual(t *testing.T) {
	data := make([]uint, 1000)
	for i := range data {
		data[i] = 7
	}
	UintDesc(data)
	for _, v := range data {
		if v != 7 {
			t.Errorf("expected all elements to be 7, got %v", data)
			break
		}
	}
}

func TestUintAsc_AlreadySorted(t *testing.T) {
	data := make([]uint, 10000)
	for i := range data {
		data[i] = uint(i)
	}
	UintAsc(data)
	for i := 0; i < len(data); i++ {
		if data[i] != uint(i) {
			t.Errorf("expected already sorted slice to remain unchanged")
			break
		}
	}
}

func TestUintDesc_AlreadySorted(t *testing.T) {
	n := 10000
	data := make([]uint, n)
	for i := range data {
		data[i] = uint(n - i - 1)
	}
	UintDesc(data)
	for i := 1; i < len(data); i++ {
		if data[i-1] < data[i] {
			t.Errorf("expected already descending slice to remain unchanged")
			break
		}
	}
}

func TestUintAsc_ReverseSorted(t *testing.T) {
	n := 10000
	data := make([]uint, n)
	for i := range data {
		data[i] = uint(n - i)
	}
	UintAsc(data)
	for i := 1; i < n; i++ {
		if data[i-1] > data[i] {
			t.Errorf("slice not sorted in ascending order")
			break
		}
	}
}

func TestUintDesc_ReverseSorted(t *testing.T) {
	n := 10000
	data := make([]uint, n)
	for i := range data {
		data[i] = uint(i)
	}
	UintDesc(data)
	for i := 1; i < n; i++ {
		if data[i-1] < data[i] {
			t.Errorf("slice not sorted in descending order")
			break
		}
	}
}

func TestUintAsc_SmallRandom(t *testing.T) {
	data := genUints(1000)
	expected := append([]uint(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	UintAsc(data)
	if !uintSlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small ascending slice")
	}
}

func TestUintDesc_SmallRandom(t *testing.T) {
	data := genUints(1000)
	expected := append([]uint(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] > expected[j] })
	UintDesc(data)
	if !uintSlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small descending slice")
	}
}

func TestUintAsc_LargeRandom(t *testing.T) {
	data := genUints(2000000)
	expected := append([]uint(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	UintAsc(data)
	if !uintSlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large slice")
	}
}

func TestUintDesc_LargeRandom(t *testing.T) {
	data := genUints(2000000)
	expected := append([]uint(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] > expected[j] })
	UintDesc(data)
	if !uintSlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large descending slice")
	}
}

func BenchmarkParsortUintAsc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Asc_Uint_"+strconv.Itoa(size), func(b *testing.B) {
			data := genUints(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]uint, len(data))
				copy(tmp, data)
				UintAsc(tmp)
			}
		})
	}
}

func BenchmarkParsortUintDesc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Desc_Uint_"+strconv.Itoa(size), func(b *testing.B) {
			data := genUints(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]uint, len(data))
				copy(tmp, data)
				UintDesc(tmp)
			}
		})
	}
}
