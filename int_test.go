package main

import (
	"sort"
	"strconv"
	"testing"
)

func TestIntAsc_EmptySlice(t *testing.T) {
	var data []int
	IntAsc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestIntDesc_EmptySlice(t *testing.T) {
	var data []int
	IntDesc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestIntAsc_SingleElement(t *testing.T) {
	data := []int{42}
	IntAsc(data)
	if data[0] != 42 {
		t.Errorf("expected [42], got %v", data)
	}
}

func TestIntDesc_SingleElement(t *testing.T) {
	data := []int{42}
	IntDesc(data)
	if data[0] != 42 {
		t.Errorf("expected [42], got %v", data)
	}
}

func TestIntAsc_AllEqual(t *testing.T) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = 7
	}
	IntAsc(data)
	for _, v := range data {
		if v != 7 {
			t.Errorf("expected all elements to be 7, got %v", data)
			break
		}
	}
}

func TestIntDesc_AllEqual(t *testing.T) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = 7
	}
	IntDesc(data)
	for _, v := range data {
		if v != 7 {
			t.Errorf("expected all elements to be 7, got %v", data)
			break
		}
	}
}

func TestIntAsc_AlreadySorted(t *testing.T) {
	data := make([]int, 10000)
	for i := range data {
		data[i] = i
	}
	IntAsc(data)
	for i := 0; i < len(data); i++ {
		if data[i] != i {
			t.Errorf("expected already sorted slice to remain unchanged")
			break
		}
	}
}

func TestIntDesc_AlreadySorted(t *testing.T) {
	n := 10000
	data := make([]int, n)
	for i := range data {
		data[i] = n - i - 1
	}
	IntDesc(data)
	for i := 1; i < len(data); i++ {
		if data[i-1] < data[i] {
			t.Errorf("expected already descending slice to remain unchanged")
			break
		}
	}
}

func TestIntAsc_ReverseSorted(t *testing.T) {
	n := 10000
	data := make([]int, n)
	for i := range data {
		data[i] = n - i
	}
	IntAsc(data)
	for i := 1; i < n; i++ {
		if data[i-1] > data[i] {
			t.Errorf("slice not sorted in ascending order")
			break
		}
	}
}

func TestIntDesc_ReverseSorted(t *testing.T) {
	n := 10000
	data := make([]int, n)
	for i := range data {
		data[i] = i
	}
	IntDesc(data)
	for i := 1; i < n; i++ {
		if data[i-1] < data[i] {
			t.Errorf("slice not sorted in descending order")
			break
		}
	}
}

func TestIntAsc_LargeRandom(t *testing.T) {
	data := genInts(2000000)
	expected := append([]int(nil), data...)
	sort.Ints(expected)
	IntAsc(data)
	if !intSlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large slice")
	}
}

func TestIntDesc_LargeRandom(t *testing.T) {
	data := genInts(2000000)
	expected := append([]int(nil), data...)
	sort.Sort(sort.Reverse(sort.IntSlice(expected)))
	IntDesc(data)
	if !intSlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large descending slice")
	}
}

func TestIntAsc_SmallRandom(t *testing.T) {
	data := genInts(1000)
	expected := append([]int(nil), data...)
	sort.Ints(expected)
	IntAsc(data)
	if !intSlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small ascending slice")
	}
}

func TestIntDesc_SmallRandom(t *testing.T) {
	data := genInts(1000)
	expected := append([]int(nil), data...)
	sort.Sort(sort.Reverse(sort.IntSlice(expected)))
	IntDesc(data)
	if !intSlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small descending slice")
	}
}

func BenchmarkParsortIntAsc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Asc_Int_"+strconv.Itoa(size), func(b *testing.B) {
			data := genInts(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]int, len(data))
				copy(tmp, data)
				IntAsc(tmp)
			}
		})
	}
}

func BenchmarkParsortIntDesc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Desc_Int_"+strconv.Itoa(size), func(b *testing.B) {
			data := genInts(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]int, len(data))
				copy(tmp, data)
				IntDesc(tmp)
			}
		})
	}
}
