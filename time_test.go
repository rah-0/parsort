package main

import (
	"sort"
	"strconv"
	"testing"
	"time"
)

func TestTimeAsc_EmptySlice(t *testing.T) {
	var data []time.Time
	TimeAsc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestTimeDesc_EmptySlice(t *testing.T) {
	var data []time.Time
	TimeDesc(data)
	if len(data) != 0 {
		t.Errorf("expected empty slice, got %v", data)
	}
}

func TestTimeAsc_SingleElement(t *testing.T) {
	now := time.Now()
	data := []time.Time{now}
	TimeAsc(data)
	if !data[0].Equal(now) {
		t.Errorf("expected [%v], got %v", now, data)
	}
}

func TestTimeDesc_SingleElement(t *testing.T) {
	now := time.Now()
	data := []time.Time{now}
	TimeDesc(data)
	if !data[0].Equal(now) {
		t.Errorf("expected [%v], got %v", now, data)
	}
}

func TestTimeAsc_AllEqual(t *testing.T) {
	t0 := time.Unix(1000000, 0)
	data := make([]time.Time, 1000)
	for i := range data {
		data[i] = t0
	}
	TimeAsc(data)
	for _, v := range data {
		if !v.Equal(t0) {
			t.Errorf("expected all elements to be %v, got %v", t0, data)
			break
		}
	}
}

func TestTimeDesc_AllEqual(t *testing.T) {
	t0 := time.Unix(1000000, 0)
	data := make([]time.Time, 1000)
	for i := range data {
		data[i] = t0
	}
	TimeDesc(data)
	for _, v := range data {
		if !v.Equal(t0) {
			t.Errorf("expected all elements to be %v, got %v", t0, data)
			break
		}
	}
}

func TestTimeAsc_AlreadySorted(t *testing.T) {
	base := time.Now()
	data := make([]time.Time, 10000)
	for i := range data {
		data[i] = base.Add(time.Duration(i) * time.Second)
	}
	TimeAsc(data)
	for i := 1; i < len(data); i++ {
		if data[i].Before(data[i-1]) {
			t.Errorf("expected already sorted slice to remain unchanged")
			break
		}
	}
}

func TestTimeDesc_AlreadySorted(t *testing.T) {
	base := time.Now()
	n := 10000
	data := make([]time.Time, n)
	for i := range data {
		data[i] = base.Add(time.Duration(n-i) * time.Second)
	}
	TimeDesc(data)
	for i := 1; i < len(data); i++ {
		if data[i].After(data[i-1]) {
			t.Errorf("expected descending slice to remain unchanged")
			break
		}
	}
}

func TestTimeAsc_ReverseSorted(t *testing.T) {
	base := time.Now()
	n := 10000
	data := make([]time.Time, n)
	for i := range data {
		data[i] = base.Add(time.Duration(n-i) * time.Second)
	}
	TimeAsc(data)
	for i := 1; i < len(data); i++ {
		if data[i].Before(data[i-1]) {
			t.Errorf("slice not sorted in ascending order")
			break
		}
	}
}

func TestTimeDesc_ReverseSorted(t *testing.T) {
	base := time.Now()
	n := 10000
	data := make([]time.Time, n)
	for i := range data {
		data[i] = base.Add(time.Duration(i) * time.Second)
	}
	TimeDesc(data)
	for i := 1; i < len(data); i++ {
		if data[i].After(data[i-1]) {
			t.Errorf("slice not sorted in descending order")
			break
		}
	}
}

func TestTimeAsc_SmallRandom(t *testing.T) {
	data := genTimes(1000)
	expected := append([]time.Time(nil), data...)
	sort.Slice(expected, func(i, j int) bool {
		return expected[i].Before(expected[j])
	})
	TimeAsc(data)
	if !timeSlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small ascending slice")
	}
}

func TestTimeDesc_SmallRandom(t *testing.T) {
	data := genTimes(1000)
	expected := append([]time.Time(nil), data...)
	sort.Slice(expected, func(i, j int) bool {
		return expected[i].After(expected[j])
	})
	TimeDesc(data)
	if !timeSlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for small descending slice")
	}
}

func TestTimeAsc_LargeRandom(t *testing.T) {
	data := genTimes(2000000)
	expected := append([]time.Time(nil), data...)
	sort.Slice(expected, func(i, j int) bool {
		return expected[i].Before(expected[j])
	})
	TimeAsc(data)
	if !timeSlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large ascending slice")
	}
}

func TestTimeDesc_LargeRandom(t *testing.T) {
	data := genTimes(2000000)
	expected := append([]time.Time(nil), data...)
	sort.Slice(expected, func(i, j int) bool {
		return expected[i].After(expected[j])
	})
	TimeDesc(data)
	if !timeSlicesEqual(data, expected) {
		t.Errorf("sorted result incorrect for large descending slice")
	}
}

func BenchmarkParsortTimeAsc(b *testing.B) {
	for _, size := range testSizes {
		b.Run("Parsort_Sort_Time_"+strconv.Itoa(size), func(b *testing.B) {
			data := genTimes(size)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tmp := make([]time.Time, len(data))
				copy(tmp, data)
				TimeAsc(tmp)
			}
		})
	}
}
