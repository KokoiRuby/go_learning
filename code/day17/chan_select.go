package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	for i := 0; i < 100; i++ {
		ch1 <- i
		ch2 <- i
	}

	wg.Add(1)
	go func() {
		for {
			select {
			case v := <-ch1:
				fmt.Println("Recv from ch1: ", v)
			case v := <-ch2:
				fmt.Println("Recv from ch2: ", v)
			}
		}
		wg.Done()
	}()
	time.Sleep(time.Second * 3)
	wg.Wait()
}
