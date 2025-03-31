package parsort

import (
	"runtime"
)

var (
	CoreCount = runtime.NumCPU()

	IntMinParallelSize   = 10000
	Int8MinParallelSize  = 5000
	Int16MinParallelSize = 5000
	Int32MinParallelSize = 5000
	Int64MinParallelSize = 5000

	UintMinParallelSize   = 5000
	Uint8MinParallelSize  = 5000
	Uint16MinParallelSize = 5000
	Uint32MinParallelSize = 5000
	Uint64MinParallelSize = 5000

	Float32MinParallelSize = 5000
	Float64MinParallelSize = 10000

	StringMinParallelSize = 10000
	StructMinParallelSize = 10000
	TimeMinParallelSize   = 5000
)
