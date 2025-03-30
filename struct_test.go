package main

import (
	"strconv"
	"testing"
)

func TestStructAsc(t *testing.T) {
	data := genPeople(50000)
	StructAsc(data, func(a, b person) bool { return a.Age < b.Age })
	if !isSortedAsc(data) {
		t.Errorf("StructAsc failed to sort correctly")
	}
}

func TestStructDesc(t *testing.T) {
	data := genPeople(50000)
	StructDesc(data, func(a, b person) bool { return a.Age < b.Age })
	if !isSortedDesc(data) {
		t.Errorf("StructDesc failed to sort correctly")
	}
}

func TestStructAsc_Small(t *testing.T) {
	data := genPeople(5)
	StructAsc(data, func(a, b person) bool { return a.Age < b.Age })
	if !isSortedAsc(data) {
		t.Errorf("Small slice sort (Asc) failed")
	}
}

func TestStructDesc_Small(t *testing.T) {
	data := genPeople(5)
	StructDesc(data, func(a, b person) bool { return a.Age < b.Age })
	if !isSortedDesc(data) {
		t.Errorf("Small slice sort (Desc) failed")
	}
}

func TestStructAsc_Empty(t *testing.T) {
	data := genPeople(0)
	StructAsc(data, func(a, b person) bool { return a.Age < b.Age })
	if len(data) != 0 {
		t.Errorf("Expected empty slice, got: %v", data)
	}
}

func TestStructAsc_Single(t *testing.T) {
	data := genPeople(1)
	expected := data[0]
	StructAsc(data, func(a, b person) bool { return a.Age < b.Age })
	if data[0] != expected {
		t.Errorf("Single element slice changed unexpectedly")
	}
}

func TestStructAsc_AlreadySorted(t *testing.T) {
	data := []person{{"A", 1}, {"B", 2}, {"C", 3}}
	StructAsc(data, func(a, b person) bool { return a.Age < b.Age })
	if !isSortedAsc(data) {
		t.Errorf("Already sorted slice failed Asc sort")
	}
}

func TestStructAsc_ReversedInput(t *testing.T) {
	data := []person{{"C", 3}, {"B", 2}, {"A", 1}}
	StructAsc(data, func(a, b person) bool { return a.Age < b.Age })
	if !isSortedAsc(data) {
		t.Errorf("Reversed input failed to sort Asc")
	}
}

func TestStructAsc_AllEqual(t *testing.T) {
	data := []person{{"A", 5}, {"B", 5}, {"C", 5}}
	StructAsc(data, func(a, b person) bool { return a.Age < b.Age })
	if !isSortedAsc(data) {
		t.Errorf("Equal elements failed to remain in order")
	}
}

func TestStructDesc_Stability(t *testing.T) {
	data := []person{{"Z", 3}, {"A", 3}, {"B", 3}}
	StructDesc(data, func(a, b person) bool { return a.Age < b.Age })
	if data[0].Name != "Z" || data[1].Name != "A" || data[2].Name != "B" {
		t.Errorf("Stable order was not preserved for equal values")
	}
}

func BenchmarkSortStruct_Arbitrary(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Arbitrary_SortStruct_"+strconv.Itoa(size), func(b *testing.B) {
			original := genPeople(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]person, len(original))
				copy(tmp, original)
				structSort(tmp, func(a, b person) bool {
					return a.Age < b.Age
				})
			}
		})
	}
}
