package main

import "fmt"

func main() {
	ch := make(chan int)

	// 1
	go func() {
		// 4
		fmt.Println("Ready to send 42")
		ch <- 42
		fmt.Println("42 is sent.")
	}()

	// 2
	fmt.Println("Waiting for value from channel...")
	// 3 Blocked
	val := <-ch
	// 5
	fmt.Println("Value received: ", val)
}
