package parsort

import (
	"sort"
	"strconv"
	"testing"
)

func TestFloat64Asc_EmptySlice(t *testing.T) {
	var data []float64
	Float64Asc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestFloat64Desc_EmptySlice(t *testing.T) {
	var data []float64
	Float64Desc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestFloat64Asc_SingleElement(t *testing.T) {
	data := []float64{42.0}
	Float64Asc(data)
	if data[0] != 42.0 {
		t.Errorf("expected [42.0], got %v", data)
	}
}

func TestFloat64Desc_SingleElement(t *testing.T) {
	data := []float64{42.0}
	Float64Desc(data)
	if data[0] != 42.0 {
		t.Errorf("expected [42.0], got %v", data)
	}
}

func TestFloat64Asc_AllEqual(t *testing.T) {
	data := make([]float64, 1000)
	for i := range data {
		data[i] = 3.14
	}
	Float64Asc(data)
	for _, v := range data {
		if v != 3.14 {
			t.Errorf("expected all elements to be 3.14, got %v", data)
			break
		}
	}
}

func TestFloat64Desc_AllEqual(t *testing.T) {
	data := make([]float64, 1000)
	for i := range data {
		data[i] = 3.14
	}
	Float64Desc(data)
	for _, v := range data {
		if v != 3.14 {
			t.Errorf("expected all elements to be 3.14, got %v", data)
			break
		}
	}
}

func TestFloat64Asc_AlreadySorted(t *testing.T) {
	data := make([]float64, 10000)
	for i := range data {
		data[i] = float64(i)
	}
	Float64Asc(data)
	for i := 0; i < len(data); i++ {
		if data[i] != float64(i) {
			t.Errorf("expected already sorted slice to remain unchanged")
			break
		}
	}
}

func TestFloat64Desc_AlreadySorted(t *testing.T) {
	n := 10000
	data := make([]float64, n)
	for i := range data {
		data[i] = float64(n - i - 1)
	}
	Float64Desc(data)
	for i := 1; i < len(data); i++ {
		if data[i-1] < data[i] {
			t.Errorf("expected already descending slice to remain unchanged")
			break
		}
	}
}

func TestFloat64Asc_ReverseSorted(t *testing.T) {
	n := 10000
	data := make([]float64, n)
	for i := range data {
		data[i] = float64(n - i)
	}
	Float64Asc(data)
	for i := 1; i < n; i++ {
		if data[i-1] > data[i] {
			t.Errorf("slice not sorted in ascending order")
			break
		}
	}
}

func TestFloat64Desc_ReverseSorted(t *testing.T) {
	n := 10000
	data := make([]float64, n)
	for i := range data {
		data[i] = float64(i)
	}
	Float64Desc(data)
	for i := 1; i < n; i++ {
		if data[i-1] < data[i] {
			t.Errorf("slice not sorted in descending order")
			break
		}
	}
}

func TestFloat64Asc_LargeRandom(t *testing.T) {
	data := genFloats(2000000)
	expected := append([]float64(nil), data...)
	sort.Float64s(expected)
	Float64Asc(data)
	if !floatSlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large slice")
	}
}

func TestFloat64Desc_LargeRandom(t *testing.T) {
	data := genFloats(2000000)
	expected := append([]float64(nil), data...)
	sort.Sort(sort.Reverse(sort.Float64Slice(expected)))
	Float64Desc(data)
	if !floatSlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large descending slice")
	}
}

func TestFloat64Asc_SmallRandom(t *testing.T) {
	data := genFloats(1000)
	expected := append([]float64(nil), data...)
	sort.Float64s(expected)
	Float64Asc(data)
	if !floatSlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small ascending slice")
	}
}

func TestFloat64Desc_SmallRandom(t *testing.T) {
	data := genFloats(1000)
	expected := append([]float64(nil), data...)
	sort.Sort(sort.Reverse(sort.Float64Slice(expected)))
	Float64Desc(data)
	if !floatSlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small descending slice")
	}
}

func BenchmarkParsortFloat64Asc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Asc_Float64_"+strconv.Itoa(size), func(b *testing.B) {
			data := genFloats(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]float64, len(data))
				copy(tmp, data)
				Float64Asc(tmp)
			}
		})
	}
}

func BenchmarkParsortFloat64Desc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Desc_Float64_"+strconv.Itoa(size), func(b *testing.B) {
			data := genFloats(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]float64, len(data))
				copy(tmp, data)
				Float64Desc(tmp)
			}
		})
	}
}
