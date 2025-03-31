package parsort

import (
	"fmt"
	"runtime"
	"sort"
	"strconv"
	"time"
)

// TuneSpecific benchmarks sequential vs. concurrent sorting for increasing data sizes,
// starting from `startSize` and increasing by `increment` each iteration.
// For each case, it runs the benchmark `runs` times.
// The loop continues indefinitely, but is intended to be stopped by logic inside
// the callback (e.g., once a performance delta threshold is met).
//
// Parameters:
// - runs:          	Number of repetitions per case for averaging
// - startSize:     	Initial slice size to test
// - increment:     	Amount to increase slice size per iteration
// - deltaThreshold: 	Intended threshold for delta evaluation (used externally in callback)
// - showOutput: 		Prints the bench results
//
// Example use case:
//   - Automatically find the minimum slice size where parallel sorting becomes faster.
//   - Tunes ParallelSize variables for each type based on real CPU performance.
func TuneSpecific(runs int, startSize int, increment int, deltaThreshold float64, showOutput bool) {
	tuner := NewTuner().SetRuns(runs)

	IntMinParallelSize = 0
	size := startSize
	stop := false
	for !stop {
		label := strconv.Itoa(size)
		sampleData := genInts(size)

		tuner.AddCase(label,
			func() {
				data := make([]int, len(sampleData))
				copy(data, sampleData)
				sort.Ints(data)
			},
			func() {
				data := make([]int, len(sampleData))
				copy(data, sampleData)
				IntAsc(data)
			},
			func(r BenchResult) {
				if r.DeltaNsPct < deltaThreshold {
					if showOutput {
						fmt.Println("Int")
						tuner.PrintResult()
					}
					IntMinParallelSize = size
					stop = true
				}
			},
		)

		tuner.Run().Reset()
		size += increment
	}

	Int8MinParallelSize = 0
	size = startSize
	stop = false
	for !stop {
		label := strconv.Itoa(size)
		sampleData := genInt8s(size)

		tuner.AddCase(label,
			func() {
				data := make([]int8, len(sampleData))
				copy(data, sampleData)
				sort.Slice(data, func(i, j int) bool {
					return data[i] < data[j]
				})
			},
			func() {
				data := make([]int8, len(sampleData))
				copy(data, sampleData)
				Int8Asc(data)
			},
			func(r BenchResult) {
				if r.DeltaNsPct < deltaThreshold {
					if showOutput {
						fmt.Println("Int8")
						tuner.PrintResult()
					}
					Int8MinParallelSize = size
					stop = true
				}
			},
		)

		tuner.Run().Reset()
		size += increment
	}

	Int16MinParallelSize = 0
	size = startSize
	stop = false
	for !stop {
		label := strconv.Itoa(size)
		sampleData := genInt16s(size)

		tuner.AddCase(label,
			func() {
				data := make([]int16, len(sampleData))
				copy(data, sampleData)
				sort.Slice(data, func(i, j int) bool {
					return data[i] < data[j]
				})
			},
			func() {
				data := make([]int16, len(sampleData))
				copy(data, sampleData)
				Int16Asc(data)
			},
			func(r BenchResult) {
				if r.DeltaNsPct < deltaThreshold {
					if showOutput {
						fmt.Println("Int16")
						tuner.PrintResult()
					}
					Int16MinParallelSize = size
					stop = true
				}
			},
		)

		tuner.Run().Reset()
		size += increment
	}

	Int32MinParallelSize = 0
	size = startSize
	stop = false
	for !stop {
		label := strconv.Itoa(size)
		sampleData := genInt32s(size)

		tuner.AddCase(label,
			func() {
				data := make([]int32, len(sampleData))
				copy(data, sampleData)
				sort.Slice(data, func(i, j int) bool {
					return data[i] < data[j]
				})
			},
			func() {
				data := make([]int32, len(sampleData))
				copy(data, sampleData)
				Int32Asc(data)
			},
			func(r BenchResult) {
				if r.DeltaNsPct < deltaThreshold {
					if showOutput {
						fmt.Println("Int32")
						tuner.PrintResult()
					}
					Int32MinParallelSize = size
					stop = true
				}
			},
		)

		tuner.Run().Reset()
		size += increment
	}

	Int64MinParallelSize = 0
	size = startSize
	stop = false
	for !stop {
		label := strconv.Itoa(size)
		sampleData := genInt64s(size)

		tuner.AddCase(label,
			func() {
				data := make([]int64, len(sampleData))
				copy(data, sampleData)
				sort.Slice(data, func(i, j int) bool {
					return data[i] < data[j]
				})
			},
			func() {
				data := make([]int64, len(sampleData))
				copy(data, sampleData)
				Int64Asc(data)
			},
			func(r BenchResult) {
				if r.DeltaNsPct < deltaThreshold {
					if showOutput {
						fmt.Println("Int64")
						tuner.PrintResult()
					}
					Int64MinParallelSize = size
					stop = true
				}
			},
		)

		tuner.Run().Reset()
		size += increment
	}

	UintMinParallelSize = 0
	size = startSize
	stop = false
	for !stop {
		label := strconv.Itoa(size)
		sampleData := genUints(size)

		tuner.AddCase(label,
			func() {
				data := make([]uint, len(sampleData))
				copy(data, sampleData)
				sort.Slice(data, func(i, j int) bool {
					return data[i] < data[j]
				})
			},
			func() {
				data := make([]uint, len(sampleData))
				copy(data, sampleData)
				UintAsc(data)
			},
			func(r BenchResult) {
				if r.DeltaNsPct < deltaThreshold {
					if showOutput {
						fmt.Println("Uint")
						tuner.PrintResult()
					}
					UintMinParallelSize = size
					stop = true
				}
			},
		)

		tuner.Run().Reset()
		size += increment
	}

	Uint8MinParallelSize = 0
	size = startSize
	stop = false
	for !stop {
		label := strconv.Itoa(size)
		sampleData := genUint8s(size)

		tuner.AddCase(label,
			func() {
				data := make([]uint8, len(sampleData))
				copy(data, sampleData)
				sort.Slice(data, func(i, j int) bool {
					return data[i] < data[j]
				})
			},
			func() {
				data := make([]uint8, len(sampleData))
				copy(data, sampleData)
				Uint8Asc(data)
			},
			func(r BenchResult) {
				if r.DeltaNsPct < deltaThreshold {
					if showOutput {
						fmt.Println("Uint8")
						tuner.PrintResult()
					}
					Uint8MinParallelSize = size
					stop = true
				}
			},
		)

		tuner.Run().Reset()
		size += increment
	}

	Uint16MinParallelSize = 0
	size = startSize
	stop = false
	for !stop {
		label := strconv.Itoa(size)
		sampleData := genUint16s(size)

		tuner.AddCase(label,
			func() {
				data := make([]uint16, len(sampleData))
				copy(data, sampleData)
				sort.Slice(data, func(i, j int) bool {
					return data[i] < data[j]
				})
			},
			func() {
				data := make([]uint16, len(sampleData))
				copy(data, sampleData)
				Uint16Asc(data)
			},
			func(r BenchResult) {
				if r.DeltaNsPct < deltaThreshold {
					if showOutput {
						fmt.Println("Uint16")
						tuner.PrintResult()
					}
					Uint16MinParallelSize = size
					stop = true
				}
			},
		)

		tuner.Run().Reset()
		size += increment
	}

	Uint32MinParallelSize = 0
	size = startSize
	stop = false
	for !stop {
		label := strconv.Itoa(size)
		sampleData := genUint32s(size)

		tuner.AddCase(label,
			func() {
				data := make([]uint32, len(sampleData))
				copy(data, sampleData)
				sort.Slice(data, func(i, j int) bool {
					return data[i] < data[j]
				})
			},
			func() {
				data := make([]uint32, len(sampleData))
				copy(data, sampleData)
				Uint32Asc(data)
			},
			func(r BenchResult) {
				if r.DeltaNsPct < deltaThreshold {
					if showOutput {
						fmt.Println("Uint32")
						tuner.PrintResult()
					}
					Uint32MinParallelSize = size
					stop = true
				}
			},
		)

		tuner.Run().Reset()
		size += increment
	}

	Uint64MinParallelSize = 0
	size = startSize
	stop = false
	for !stop {
		label := strconv.Itoa(size)
		sampleData := genUint64s(size)

		tuner.AddCase(label,
			func() {
				data := make([]uint64, len(sampleData))
				copy(data, sampleData)
				sort.Slice(data, func(i, j int) bool {
					return data[i] < data[j]
				})
			},
			func() {
				data := make([]uint64, len(sampleData))
				copy(data, sampleData)
				Uint64Asc(data)
			},
			func(r BenchResult) {
				if r.DeltaNsPct < deltaThreshold {
					if showOutput {
						fmt.Println("Uint64")
						tuner.PrintResult()
					}
					Uint64MinParallelSize = size
					stop = true
				}
			},
		)

		tuner.Run().Reset()
		size += increment
	}

	Float32MinParallelSize = 0
	size = startSize
	stop = false
	for !stop {
		label := strconv.Itoa(size)
		sampleData := genFloat32s(size)

		tuner.AddCase(label,
			func() {
				data := make([]float32, len(sampleData))
				copy(data, sampleData)
				sort.Slice(data, func(i, j int) bool {
					return data[i] < data[j]
				})
			},
			func() {
				data := make([]float32, len(sampleData))
				copy(data, sampleData)
				Float32Asc(data)
			},
			func(r BenchResult) {
				if r.DeltaNsPct < deltaThreshold {
					if showOutput {
						fmt.Println("Float32")
						tuner.PrintResult()
					}
					Float32MinParallelSize = size
					stop = true
				}
			},
		)

		tuner.Run().Reset()
		size += increment
	}

	Float64MinParallelSize = 0
	size = startSize
	stop = false
	for !stop {
		label := strconv.Itoa(size)
		sampleData := genFloats(size)

		tuner.AddCase(label,
			func() {
				data := make([]float64, len(sampleData))
				copy(data, sampleData)
				sort.Float64s(data)
			},
			func() {
				data := make([]float64, len(sampleData))
				copy(data, sampleData)
				Float64Asc(data)
			},
			func(r BenchResult) {
				if r.DeltaNsPct < deltaThreshold {
					if showOutput {
						fmt.Println("Float64")
						tuner.PrintResult()
					}
					Float64MinParallelSize = size
					stop = true
				}
			},
		)

		tuner.Run().Reset()
		size += increment
	}

	StringMinParallelSize = 0
	size = startSize
	stop = false
	for !stop {
		label := strconv.Itoa(size)
		sampleData := genStrings(size)

		tuner.AddCase(label,
			func() {
				data := make([]string, len(sampleData))
				copy(data, sampleData)
				sort.Strings(data)
			},
			func() {
				data := make([]string, len(sampleData))
				copy(data, sampleData)
				StringAsc(data)
			},
			func(r BenchResult) {
				if r.DeltaNsPct < deltaThreshold {
					if showOutput {
						fmt.Println("String")
						tuner.PrintResult()
					}
					StringMinParallelSize = size
					stop = true
				}
			},
		)

		tuner.Run().Reset()
		size += increment
	}

	TimeMinParallelSize = 0
	size = startSize
	stop = false
	for !stop {
		label := strconv.Itoa(size)
		sampleData := genTimes(size)

		tuner.AddCase(label,
			func() {
				data := make([]time.Time, len(sampleData))
				copy(data, sampleData)
				sort.Slice(data, func(i, j int) bool {
					return data[i].Before(data[j])
				})
			},
			func() {
				data := make([]time.Time, len(sampleData))
				copy(data, sampleData)
				TimeAsc(data)
			},
			func(r BenchResult) {
				if r.DeltaNsPct < deltaThreshold {
					if showOutput {
						fmt.Println("Time")
						tuner.PrintResult()
					}
					TimeMinParallelSize = size
					stop = true
				}
			},
		)

		tuner.Run().Reset()
		size += increment
	}

	StructMinParallelSize = 0
	size = startSize
	stop = false
	for !stop {
		label := strconv.Itoa(size)
		sampleData := genPeople(size)

		tuner.AddCase(label,
			func() {
				data := make([]person, len(sampleData))
				copy(data, sampleData)
				sort.Sort(byAge(data))
			},
			func() {
				data := make([]person, len(sampleData))
				copy(data, sampleData)
				StructAsc(data, func(a, b person) bool { return a.Age < b.Age })
			},
			func(r BenchResult) {
				if r.DeltaNsPct < deltaThreshold {
					if showOutput {
						fmt.Println("Struct")
						tuner.PrintResult()
					}
					StructMinParallelSize = size
					stop = true
				}
			},
		)

		tuner.Run().Reset()
		size += increment
	}
}

func Tune() {
	TuneSpecific(100, 1000, 1000, -10, false)
}

type Tuner struct {
	runsPerCase int
	cases       []benchCase
	results     []BenchResult
}

type BenchResult struct {
	Label      string
	SeqTimeNs  int64
	ConTimeNs  int64
	SeqBytes   int64
	ConBytes   int64
	DeltaNsPct float64
	DeltaBPct  float64
}

type benchCase struct {
	Label          string
	FuncSequential func()
	FuncConcurrent func()
	Callback       func(BenchResult)
}

func NewTuner() *Tuner {
	return &Tuner{
		runsPerCase: 5,
	}
}

func (x *Tuner) SetRuns(n int) *Tuner {
	x.runsPerCase = n
	return x
}

func (x *Tuner) AddCase(label string, seq, con func(), cb func(BenchResult)) *Tuner {
	x.cases = append(x.cases, benchCase{
		Label:          label,
		FuncSequential: seq,
		FuncConcurrent: con,
		Callback:       cb,
	})
	return x
}

func (x *Tuner) Reset() {
	x.cases = []benchCase{}
	x.results = []BenchResult{}
}

func (x *Tuner) Run() *Tuner {
	for _, c := range x.cases {
		var seqTotalNs, conTotalNs, seqTotalB, conTotalB int64

		for i := 0; i < x.runsPerCase; i++ {
			seqAlloc := x.measureAlloc(c.FuncSequential)
			seqTotalNs += seqAlloc.timeNs
			seqTotalB += seqAlloc.bytes

			conAlloc := x.measureAlloc(c.FuncConcurrent)
			conTotalNs += conAlloc.timeNs
			conTotalB += conAlloc.bytes
		}

		seqAvgNs := seqTotalNs / int64(x.runsPerCase)
		conAvgNs := conTotalNs / int64(x.runsPerCase)
		seqAvgB := seqTotalB / int64(x.runsPerCase)
		conAvgB := conTotalB / int64(x.runsPerCase)

		deltaNsPct := float64(conAvgNs-seqAvgNs) / float64(seqAvgNs) * 100
		deltaBPct := float64(conAvgB-seqAvgB) / float64(seqAvgB) * 100

		result := BenchResult{
			Label:      c.Label,
			SeqTimeNs:  seqAvgNs,
			ConTimeNs:  conAvgNs,
			SeqBytes:   seqAvgB,
			ConBytes:   conAvgB,
			DeltaNsPct: deltaNsPct,
			DeltaBPct:  deltaBPct,
		}
		x.results = append(x.results, result)
		if c.Callback != nil {
			c.Callback(result)
		}
	}
	return x
}

type allocResult struct {
	timeNs int64
	bytes  int64
}

func (x *Tuner) measureAlloc(fn func()) allocResult {
	var memBefore, memAfter runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&memBefore)
	start := time.Now()
	fn()
	dur := time.Since(start).Nanoseconds()
	runtime.ReadMemStats(&memAfter)
	return allocResult{
		timeNs: dur,
		bytes:  int64(memAfter.TotalAlloc - memBefore.TotalAlloc),
	}
}

func (x *Tuner) PrintResult() *Tuner {
	titles := []string{"Label", "Sequential ns/op", "Concurrent ns/op", "Sequential B/op", "Concurrent B/op", "Delta ns/op", "Delta B/op"}
	maxLabel := len(titles[0])
	for _, r := range x.results {
		maxLabel = x.max(maxLabel, len(r.Label))
	}
	maxNs := len(titles[1])
	maxB := len(titles[3])
	for _, r := range x.results {
		maxNs = x.max(maxNs, len(x.HumanNs(r.SeqTimeNs)))
		maxNs = x.max(maxNs, len(x.HumanNs(r.ConTimeNs)))
		maxB = x.max(maxB, len(x.HumanBytes(r.SeqBytes)))
		maxB = x.max(maxB, len(x.HumanBytes(r.ConBytes)))
	}

	fmt.Printf("| %-*s | %s | %s | %s | %s | %s | %s |\n",
		maxLabel, titles[0],
		x.padCenter(titles[1], maxNs),
		x.padCenter(titles[2], maxNs),
		x.padCenter(titles[3], maxB),
		x.padCenter(titles[4], maxB),
		x.padCenter(titles[5], 12),
		x.padCenter(titles[6], 11))
	fmt.Printf("|-%s-|-%s-|-%s-|-%s-|-%s-|-%s-|-%s-|\n",
		x.repeat("-", maxLabel), x.repeat("-", maxNs), x.repeat("-", maxNs), x.repeat("-", maxB), x.repeat("-", maxB), x.repeat("-", 12), x.repeat("-", 11))

	for _, r := range x.results {
		deltaNs := fmt.Sprintf("%+.2f%%", r.DeltaNsPct)
		deltaB := fmt.Sprintf("%+.2f%%", r.DeltaBPct)
		fmt.Printf("| %-*s | %*s | %*s | %*s | %*s | %12s | %11s |\n",
			maxLabel, r.Label,
			maxNs, x.HumanNs(r.SeqTimeNs),
			maxNs, x.HumanNs(r.ConTimeNs),
			maxB, x.HumanBytes(r.SeqBytes),
			maxB, x.HumanBytes(r.ConBytes),
			deltaNs,
			deltaB)
	}
	fmt.Println()
	return x
}

func (x *Tuner) HumanNs(ns int64) string {
	out := ""
	if ns >= 3600000000000 {
		out += fmt.Sprintf("%dh ", ns/3600000000000)
		ns = ns % 3600000000000
	}
	if ns >= 60000000000 {
		out += fmt.Sprintf("%dm ", ns/60000000000)
		ns = ns % 60000000000
	}
	if ns >= 1000000000 {
		out += fmt.Sprintf("%ds ", ns/1000000000)
		ns = ns % 1000000000
	}
	if ns >= 1000000 {
		out += fmt.Sprintf("%dms ", ns/1000000)
		ns = ns % 1000000
	}
	if ns >= 1000 {
		out += fmt.Sprintf("%dµs ", ns/1000)
		ns = ns % 1000
	}
	if ns > 0 || out == "" {
		out += fmt.Sprintf("%dns", ns)
	}
	return out
}

func (x *Tuner) HumanBytes(b int64) string {
	out := ""
	if b >= 1073741824 {
		out += fmt.Sprintf("%dGB ", b/1073741824)
		b = b % 1073741824
	}
	if b >= 1048576 {
		out += fmt.Sprintf("%dMB ", b/1048576)
		b = b % 1048576
	}
	if b >= 1024 {
		out += fmt.Sprintf("%dKB ", b/1024)
		b = b % 1024
	}
	if b > 0 || out == "" {
		out += fmt.Sprintf("%dB", b)
	}
	return out
}

func (x *Tuner) repeat(s string, count int) string {
	res := ""
	for i := 0; i < count; i++ {
		res += s
	}
	return res
}

func (x *Tuner) padCenter(s string, width int) string {
	pad := width - len(s)
	if pad <= 0 {
		return s
	}
	left := pad / 2
	right := pad - left
	return x.repeat(" ", left) + s + x.repeat(" ", right)
}

func (x *Tuner) max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
