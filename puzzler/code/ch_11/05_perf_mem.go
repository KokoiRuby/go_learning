package main

import (
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	// set profiling rate, default 512KB
	// runtime.MemProfileRate = 256 * 1024

	f, err := os.Create("mem.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// active gc
	runtime.GC()

	// write to prof
	if err := pprof.WriteHeapProfile(f); err != nil {
		panic(err)
	}
}
