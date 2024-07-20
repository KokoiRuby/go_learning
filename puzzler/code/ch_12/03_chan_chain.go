package main

import (
	"fmt"
	"time"
)

func stage1(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			// Stage 1 processing
			result := num * 2
			out <- result
		}
	}()
	return out
}

func stage2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			// Stage 2 processing
			result := num + 1
			out <- result
		}
	}()
	return out
}

func main() {
	// Create input channel
	input := make(chan int)

	// Start the goroutine chain
	output := stage2(stage1(input))

	// Send input values to the chain
	go func() {
		defer close(input)
		for i := 1; i <= 5; i++ {
			input <- i
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Receive and print the final output from the chain
	for result := range output {
		fmt.Println("Output:", result)
	}
}
