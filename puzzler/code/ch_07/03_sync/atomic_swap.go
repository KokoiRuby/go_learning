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
		oldValue := atomic.SwapInt32(&value, 100)
		fmt.Println("Old Value:", oldValue)
	}()

	wg.Wait()
	fmt.Println("Final Value:", value)
}
