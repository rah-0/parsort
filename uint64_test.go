package parsort

import (
	"sort"
	"strconv"
	"testing"
)

func TestUint64Asc_EmptySlice(t *testing.T) {
	var data []uint64
	Uint64Asc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestUint64Desc_EmptySlice(t *testing.T) {
	var data []uint64
	Uint64Desc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestUint64Asc_SingleElement(t *testing.T) {
	data := []uint64{42}
	Uint64Asc(data)
	if data[0] != 42 {
		t.Errorf("expected [42], got %v", data)
	}
}

func TestUint64Desc_SingleElement(t *testing.T) {
	data := []uint64{42}
	Uint64Desc(data)
	if data[0] != 42 {
		t.Errorf("expected [42], got %v", data)
	}
}

func TestUint64Asc_AllEqual(t *testing.T) {
	data := make([]uint64, 1000)
	for i := range data {
		data[i] = 7
	}
	Uint64Asc(data)
	for _, v := range data {
		if v != 7 {
			t.Errorf("expected all elements to be 7, got %v", data)
			break
		}
	}
}

func TestUint64Desc_AllEqual(t *testing.T) {
	data := make([]uint64, 1000)
	for i := range data {
		data[i] = 7
	}
	Uint64Desc(data)
	for _, v := range data {
		if v != 7 {
			t.Errorf("expected all elements to be 7, got %v", data)
			break
		}
	}
}

func TestUint64Asc_AlreadySorted(t *testing.T) {
	data := make([]uint64, 10000)
	for i := range data {
		data[i] = uint64(i)
	}
	Uint64Asc(data)
	for i := 0; i < len(data); i++ {
		if data[i] != uint64(i) {
			t.Errorf("expected already sorted slice to remain unchanged")
			break
		}
	}
}

func TestUint64Desc_AlreadySorted(t *testing.T) {
	n := 10000
	data := make([]uint64, n)
	for i := range data {
		data[i] = uint64(n - i - 1)
	}
	Uint64Desc(data)
	for i := 1; i < len(data); i++ {
		if data[i-1] < data[i] {
			t.Errorf("expected descending slice to remain unchanged")
			break
		}
	}
}

func TestUint64Asc_ReverseSorted(t *testing.T) {
	n := 10000
	data := make([]uint64, n)
	for i := range data {
		data[i] = uint64(n - i)
	}
	Uint64Asc(data)
	for i := 1; i < n; i++ {
		if data[i-1] > data[i] {
			t.Errorf("slice not sorted in ascending order")
			break
		}
	}
}

func TestUint64Desc_ReverseSorted(t *testing.T) {
	n := 10000
	data := make([]uint64, n)
	for i := range data {
		data[i] = uint64(i)
	}
	Uint64Desc(data)
	for i := 1; i < n; i++ {
		if data[i-1] < data[i] {
			t.Errorf("slice not sorted in descending order")
			break
		}
	}
}

func TestUint64Asc_SmallRandom(t *testing.T) {
	data := genUint64s(1000)
	expected := append([]uint64(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Uint64Asc(data)
	if !uint64SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small ascending slice")
	}
}

func TestUint64Desc_SmallRandom(t *testing.T) {
	data := genUint64s(1000)
	expected := append([]uint64(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] > expected[j] })
	Uint64Desc(data)
	if !uint64SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small descending slice")
	}
}

func TestUint64Asc_LargeRandom(t *testing.T) {
	data := genUint64s(2000000)
	expected := append([]uint64(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Uint64Asc(data)
	if !uint64SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large slice")
	}
}

func TestUint64Desc_LargeRandom(t *testing.T) {
	data := genUint64s(2000000)
	expected := append([]uint64(nil), data...)
	sort.Slice(expected, func(i, j int) bool { return expected[i] > expected[j] })
	Uint64Desc(data)
	if !uint64SlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large descending slice")
	}
}

func BenchmarkParsortUint64Asc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Asc_Uint64_"+strconv.Itoa(size), func(b *testing.B) {
			data := genUint64s(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]uint64, len(data))
				copy(tmp, data)
				Uint64Asc(tmp)
			}
		})
	}
}

func BenchmarkParsortUint64Desc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Desc_Uint64_"+strconv.Itoa(size), func(b *testing.B) {
			data := genUint64s(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]uint64, len(data))
				copy(tmp, data)
				Uint64Desc(tmp)
			}
		})
	}
}
