package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(500 * time.Millisecond)
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(700 * time.Millisecond)
			ch2 <- i
		}
		close(ch2)
	}()

OuterLoop:
	for {
		select {
		case num, ok := <-ch1:
			if !ok {
				fmt.Println("ch1 is closed. Switching off.")
				ch1 = nil
			} else {
				fmt.Println("Received from ch1:", num)
			}
		case num, ok := <-ch2:
			if !ok {
				fmt.Println("ch2 is closed. Exiting loop.")
				ch2 = nil
			} else {
				fmt.Println("Received from ch2:", num)
			}
		case <-time.After(1 * time.Second):
			fmt.Println("Waiting...")
			if ch1 == nil && ch2 == nil {
				fmt.Println("Both channels are closed. Exiting loop.")
				break OuterLoop
			}
		}
	}

	fmt.Println("Main function is exiting.")
}
