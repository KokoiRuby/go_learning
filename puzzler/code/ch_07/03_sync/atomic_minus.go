package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int32 = 100
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, -1) // 原子减法
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter) // 输出: Final Counter: 90
}
