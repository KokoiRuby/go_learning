package main

import (
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

func main() {
	// sample rate
	runtime.SetBlockProfileRate(1)

	f, err := os.Create("block.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// sim block
	var mu sync.Mutex
	mu.Lock()
	go func() {
		time.Sleep(100 * time.Millisecond)
		mu.Unlock()
	}()
	mu.Lock() // trigger block

	pprof.Lookup("block").WriteTo(f, 0)
}
