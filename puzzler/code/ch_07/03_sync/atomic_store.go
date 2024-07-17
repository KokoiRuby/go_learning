package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var value int32
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		atomic.StoreInt32(&value, 42)
		fmt.Println("Value stored as 42")
	}()

	wg.Wait()
	fmt.Println("Final Value:", value)
}
