package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var numCalcsCreated int32

func createBuffer() interface{} {
	// atomic to increment
	atomic.AddInt32(&numCalcsCreated, 1)
	buffer := make([]byte, 1024)
	return &buffer
}

func main() {
	// create pool
	bufferPool := &sync.Pool{
		New: createBuffer,
	}

	// concurr
	numWorkers := 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			// get from
			buffer := bufferPool.Get()
			// buffer := createBuffer()
			_ = buffer.(*[]byte)
			// put back
			defer bufferPool.Put(buffer)
		}()
	}
	wg.Wait()
	fmt.Printf("%d buffer objects were created.\n", numCalcsCreated)
}
