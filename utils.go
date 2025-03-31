package parsort

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

var testSizes = []int{1000, 10000, 100000, 1000000, 10000000}

func genInts(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = rand.Int()
	}
	return a
}

func intSlicesEqual(a, b []int) bool {
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

func genInt8s(n int) []int8 {
	a := make([]int8, n)
	for i := range a {
		a[i] = int8(rand.Intn(256) - 128)
	}
	return a
}

func int8SlicesEqual(a, b []int8) bool {
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

func genInt16s(n int) []int16 {
	a := make([]int16, n)
	for i := range a {
		a[i] = int16(rand.Intn(65536) - 32768)
	}
	return a
}

func int16SlicesEqual(a, b []int16) bool {
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

func genInt32s(n int) []int32 {
	a := make([]int32, n)
	for i := range a {
		a[i] = rand.Int31()
	}
	return a
}

func int32SlicesEqual(a, b []int32) bool {
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

func genInt64s(n int) []int64 {
	a := make([]int64, n)
	for i := range a {
		a[i] = rand.Int63()
	}
	return a
}

func int64SlicesEqual(a, b []int64) bool {
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

func genUints(n int) []uint {
	a := make([]uint, n)
	for i := range a {
		a[i] = uint(rand.Uint32()) // 32-bit range is safe for most platforms
	}
	return a
}

func uintSlicesEqual(a, b []uint) bool {
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

func genUint8s(n int) []uint8 {
	a := make([]uint8, n)
	for i := range a {
		a[i] = uint8(rand.Intn(256))
	}
	return a
}

func uint8SlicesEqual(a, b []uint8) bool {
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

func genUint16s(n int) []uint16 {
	a := make([]uint16, n)
	for i := range a {
		a[i] = uint16(rand.Intn(65536))
	}
	return a
}

func uint16SlicesEqual(a, b []uint16) bool {
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

func genUint32s(n int) []uint32 {
	a := make([]uint32, n)
	for i := range a {
		a[i] = rand.Uint32()
	}
	return a
}

func uint32SlicesEqual(a, b []uint32) bool {
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

func genUint64s(n int) []uint64 {
	a := make([]uint64, n)
	for i := range a {
		a[i] = rand.Uint64()
	}
	return a
}

func uint64SlicesEqual(a, b []uint64) bool {
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

func genFloat32s(n int) []float32 {
	a := make([]float32, n)
	for i := range a {
		a[i] = rand.Float32()
	}
	return a
}

func float32SlicesEqual(a, b []float32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Abs(float64(a[i]-b[i])) > 1e-5 {
			return false
		}
	}
	return true
}

func genFloats(n int) []float64 {
	a := make([]float64, n)
	for i := range a {
		a[i] = rand.Float64()
	}
	return a
}

func floatSlicesEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Abs(a[i]-b[i]) > 1e-9 {
			return false
		}
	}
	return true
}

func genStrings(n int) []string {
	a := make([]string, n)
	for i := range a {
		a[i] = "str" + strconv.Itoa(rand.Intn(1000000))
	}
	return a
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

func genTimes(n int) []time.Time {
	a := make([]time.Time, n)
	base := time.Now()
	for i := range a {
		a[i] = base.Add(time.Duration(rand.Int63n(1e9)))
	}
	return a
}

func timeSlicesEqual(a, b []time.Time) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !a[i].Equal(b[i]) {
			return false
		}
	}
	return true
}

type person struct {
	Name string
	Age  int
}

func genPeople(n int) []person {
	a := make([]person, n)
	for i := range a {
		a[i] = person{
			Name: fmt.Sprintf("Name%d", rand.Intn(1000000)),
			Age:  rand.Intn(100),
		}
	}
	return a
}

type byAge []person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func isSortedAsc(data []person) bool {
	for i := 1; i < len(data); i++ {
		if data[i-1].Age > data[i].Age {
			return false
		}
	}
	return true
}

func isSortedDesc(data []person) bool {
	for i := 1; i < len(data); i++ {
		if data[i-1].Age < data[i].Age {
			return false
		}
	}
	return true
}
