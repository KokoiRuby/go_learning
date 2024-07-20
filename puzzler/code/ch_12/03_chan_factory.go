package main

import (
	"fmt"
)

func main() {
	stream := pump()
	go suck(stream)
}

// factory to gen chan
func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
