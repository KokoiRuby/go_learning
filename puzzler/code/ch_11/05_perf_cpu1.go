package main

import (
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	// make cpu work
	for i := 0; i < 1000000; i++ {
		_ = i * i
	}

	// sleep for a while to capture
	time.Sleep(1 * time.Second)
}
