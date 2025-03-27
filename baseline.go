package parsort

import (
	"math/rand"
	"strconv"
	"time"
)

func genInts(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = rand.Int()
	}
	return a
}

func genFloats(n int) []float64 {
	a := make([]float64, n)
	for i := range a {
		a[i] = rand.Float64()
	}
	return a
}

func genTimes(n int) []time.Time {
	a := make([]time.Time, n)
	base := time.Now()
	for i := range a {
		a[i] = base.Add(time.Duration(rand.Int63n(1e9)))
	}
	return a
}

func genStrings(n int) []string {
	a := make([]string, n)
	for i := range a {
		a[i] = "str" + strconv.Itoa(rand.Intn(1_000_000))
	}
	return a
}
