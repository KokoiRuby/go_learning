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
		if atomic.CompareAndSwapInt32(&value, 42, 100) {
			fmt.Println("Value was 42, changed to 100")
		}
	}()

	wg.Wait()
	fmt.Println("Final Value:", value)
}
