package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func recvChan(ch <-chan int) {
	defer wg.Done()
	for {
		num := <-ch
		if num == 0 {
			break
		}
		fmt.Println(num)
		time.Sleep(time.Second * 1)
	}
}

func sendChan(ch chan<- int) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(time.Second * 1)
	}
	close(ch)
}

func main() {
	ch := make(chan int, 10)
	wg.Add(1)
	go sendChan(ch)
	wg.Add(1)
	go recvChan(ch)
	wg.Wait()
}
