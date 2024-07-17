package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var value int32 = 42
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		loadedValue := atomic.LoadInt32(&value)
		fmt.Println("Loaded Value:", loadedValue)
	}()

	wg.Wait()
}
