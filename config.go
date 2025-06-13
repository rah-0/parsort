package parsort

import (
	"runtime"
)

var (
	// CoreCount determines the number of parallel operations to perform.
	// By default, it uses all available CPU cores.
	CoreCount = runtime.NumCPU()

	// MinParallelSize variables define the threshold at which parallel sorting becomes
	// more efficient than sequential sorting for each data type.
	// For slices smaller than these values, sequential sorting is used instead.
	// These values can be fine-tuned for specific hardware using the Tune() function.
	
	// Integer type thresholds
	IntMinParallelSize   = 10000
	Int8MinParallelSize  = 5000
	Int16MinParallelSize = 5000
	Int32MinParallelSize = 5000
	Int64MinParallelSize = 5000

	// Unsigned integer type thresholds
	UintMinParallelSize   = 5000
	Uint8MinParallelSize  = 5000
	Uint16MinParallelSize = 5000
	Uint32MinParallelSize = 5000
	Uint64MinParallelSize = 5000

	// Floating point type thresholds
	Float32MinParallelSize = 5000
	Float64MinParallelSize = 10000

	// Other type thresholds
	StringMinParallelSize = 10000
	StructMinParallelSize = 10000
	TimeMinParallelSize   = 5000
)
