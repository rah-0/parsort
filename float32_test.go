package main

import (
	"sort"
	"strconv"
	"testing"
)

func TestFloat32Asc_EmptySlice(t *testing.T) {
	var data []float32
	Float32Asc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestFloat32Desc_EmptySlice(t *testing.T) {
	var data []float32
	Float32Desc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestFloat32Asc_SingleElement(t *testing.T) {
	data := []float32{3.14}
	Float32Asc(data)
	if data[0] != 3.14 {
		t.Errorf("expected [3.14], got %v", data)
	}
}

func TestFloat32Desc_SingleElement(t *testing.T) {
	data := []float32{3.14}
	Float32Desc(data)
	if data[0] != 3.14 {
		t.Errorf("expected [3.14], got %v", data)
	}
}

func TestFloat32Asc_AllEqual(t *testing.T) {
	data := make([]float32, 1000)
	for i := range data {
		data[i] = 7.5
	}
	Float32Asc(data)
	for _, v := range data {
		if v != 7.5 {
			t.Errorf("expected all elements to be 7.5, got %v", data)
			break
		}
	}
}

func TestFloat32Desc_AllEqual(t *testing.T) {
	data := make([]float32, 1000)
	for i := range data {
		data[i] = 7.5
	}
	Float32Desc(data)
	for _, v := range data {
		if v != 7.5 {
			t.Errorf("expected all elements to be 7.5, got %v", data)
			break
		}
	}
}

func TestFloat32Asc_AlreadySorted(t *testing.T) {
	data := make([]float32, 10000)
	for i := range data {
		data[i] = float32(i)
	}
	Float32Asc(data)
	for i := 1; i < len(data); i++ {
		if data[i-1] > data[i] {
			t.Errorf("slice not sorted in ascending order")
			break
		}
	}
}

func TestFloat32Desc_AlreadySorted(t *testing.T) {
	n := 10000
	data := make([]float32, n)
	for i := range data {
		data[i] = float32(n - i - 1)
	}
	Float32Desc(data)
	for i := 1; i < len(data); i++ {
		if data[i-1] < data[i] {
			t.Errorf("slice not sorted in descending order")
			break
		}
	}
}

func TestFloat32Asc_ReverseSorted(t *testing.T) {
	n := 10000
	data := make([]float32, n)
	for i := range data {
		data[i] = float32(n - i)
	}
	Float32Asc(data)
	for i := 1; i < len(data); i++ {
		if data[i-1] > data[i] {
			t.Errorf("slice not sorted in ascending order")
			break
		}
	}
}

func TestFloat32Desc_ReverseSorted(t *testing.T) {
	n := 10000
	data := make([]float32, n)
	for i := range data {
		data[i] = float32(i)
	}
	Float32Desc(data)
	for i := 1; i < len(data); i++ {
		if data[i-1] < data[i] {
			t.Errorf("slice not sorted in descending order")
			break
		}
	}
}

func TestFloat32Asc_SmallRandom(t *testing.T) {
	data := genFloat32s(1000)
	expected := append([]float32(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Float32Asc(data)
	if !float32SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small ascending slice")
	}
}

func TestFloat32Desc_SmallRandom(t *testing.T) {
	data := genFloat32s(1000)
	expected := append([]float32(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] > expected[j] })
	Float32Desc(data)
	if !float32SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small descending slice")
	}
}

func TestFloat32Asc_LargeRandom(t *testing.T) {
	data := genFloat32s(2000000)
	expected := append([]float32(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Float32Asc(data)
	if !float32SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large slice")
	}
}

func TestFloat32Desc_LargeRandom(t *testing.T) {
	data := genFloat32s(2000000)
	expected := append([]float32(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] > expected[j] })
	Float32Desc(data)
	if !float32SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large descending slice")
	}
}

func BenchmarkParsortFloat32Asc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Asc_Float32_"+strconv.Itoa(size), func(b *testing.B) {
			data := genFloat32s(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]float32, len(data))
				copy(tmp, data)
				Float32Asc(tmp)
			}
		})
	}
}

func BenchmarkParsortFloat32Desc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Desc_Float32_"+strconv.Itoa(size), func(b *testing.B) {
			data := genFloat32s(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]float32, len(data))
				copy(tmp, data)
				Float32Desc(tmp)
			}
		})
	}
}
