package parsort

import (
	"sort"
	"strconv"
	"testing"
)

func TestStringAsc_EmptySlice(t *testing.T) {
	var data []string
	StringAsc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestStringDesc_EmptySlice(t *testing.T) {
	var data []string
	StringDesc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestStringAsc_SingleElement(t *testing.T) {
	data := []string{"abc"}
	StringAsc(data)
	if data[0] != "abc" {
		t.Errorf("expected [\"abc\"], got %v", data)
	}
}

func TestStringDesc_SingleElement(t *testing.T) {
	data := []string{"abc"}
	StringDesc(data)
	if data[0] != "abc" {
		t.Errorf("expected [\"abc\"], got %v", data)
	}
}

func TestStringAsc_AllEqual(t *testing.T) {
	data := make([]string, 1000)
	for i := range data {
		data[i] = "same"
	}
	StringAsc(data)
	for _, v := range data {
		if v != "same" {
			t.Errorf("expected all elements to be \"same\", got %v", data)
			break
		}
	}
}

func TestStringDesc_AllEqual(t *testing.T) {
	data := make([]string, 1000)
	for i := range data {
		data[i] = "same"
	}
	StringDesc(data)
	for _, v := range data {
		if v != "same" {
			t.Errorf("expected all elements to be \"same\", got %v", data)
			break
		}
	}
}

func TestStringAsc_AlreadySorted(t *testing.T) {
	data := []string{"a", "b", "c", "d", "e"}
	StringAsc(data)
	for i := 1; i < len(data); i++ {
		if data[i-1] > data[i] {
			t.Errorf("expected sorted slice to remain unchanged")
			break
		}
	}
}

func TestStringDesc_AlreadySorted(t *testing.T) {
	data := []string{"z", "y", "x", "w", "v"}
	StringDesc(data)
	for i := 1; i < len(data); i++ {
		if data[i-1] < data[i] {
			t.Errorf("expected descending slice to remain unchanged")
			break
		}
	}
}

func TestStringAsc_ReverseSorted(t *testing.T) {
	data := []string{"z", "y", "x", "w", "v"}
	StringAsc(data)
	for i := 1; i < len(data); i++ {
		if data[i-1] > data[i] {
			t.Errorf("slice not sorted in ascending order")
			break
		}
	}
}

func TestStringDesc_ReverseSorted(t *testing.T) {
	data := []string{"a", "b", "c", "d", "e"}
	StringDesc(data)
	for i := 1; i < len(data); i++ {
		if data[i-1] < data[i] {
			t.Errorf("slice not sorted in descending order")
			break
		}
	}
}

func TestStringAsc_SmallRandom(t *testing.T) {
	data := genStrings(1000)
	expected := append([]string(nil), data...)
	sort.Strings(expected)
	StringAsc(data)
	if !stringSlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small ascending slice")
	}
}

func TestStringDesc_SmallRandom(t *testing.T) {
	data := genStrings(1000)
	expected := append([]string(nil), data...)
	sort.Sort(sort.Reverse(sort.StringSlice(expected)))
	StringDesc(data)
	if !stringSlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small descending slice")
	}
}

func TestStringAsc_LargeRandom(t *testing.T) {
	data := genStrings(2_000_000)
	expected := append([]string(nil), data...)
	sort.Strings(expected)
	StringAsc(data)
	if !stringSlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large ascending slice")
	}
}

func TestStringDesc_LargeRandom(t *testing.T) {
	data := genStrings(2_000_000)
	expected := append([]string(nil), data...)
	sort.Sort(sort.Reverse(sort.StringSlice(expected)))
	StringDesc(data)
	if !stringSlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large descending slice")
	}
}

func stringSlicesEqual(a, b []string) bool {
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

func BenchmarkParsortStringAsc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Asc_String_"+strconv.Itoa(size), func(b *testing.B) {
			data := genStrings(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]string, len(data))
				copy(tmp, data)
				StringAsc(tmp)
			}
		})
	}
}

func BenchmarkParsortStringDesc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Desc_String_"+strconv.Itoa(size), func(b *testing.B) {
			data := genStrings(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]string, len(data))
				copy(tmp, data)
				StringDesc(tmp)
			}
		})
	}
}
