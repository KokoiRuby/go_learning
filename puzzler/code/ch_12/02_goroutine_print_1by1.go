package main

import (
	"fmt"
	"sync"
)

func printNumbers(wg *sync.WaitGroup, in, out chan struct{}, start, end, step int) {
	defer wg.Done()
	for i := start; i <= end; i += step {
		<-in // unblock current chan
		fmt.Println(i)
		out <- struct{}{} // notify next goroutine
	}
}

func main() {
	ch1 := make(chan struct{}, 1) // ++ buffer to avoid deadlock
	ch2 := make(chan struct{}, 1)
	ch3 := make(chan struct{}, 1)

	var wg sync.WaitGroup
	wg.Add(3)

	go printNumbers(&wg, ch1, ch2, 1, 100, 3) // 1, 4, 7...
	go printNumbers(&wg, ch2, ch3, 2, 100, 3) // 2, 5, 8...
	go printNumbers(&wg, ch3, ch1, 3, 100, 3) // 3, 6, 9...

	ch1 <- struct{}{}

	wg.Wait()
}
