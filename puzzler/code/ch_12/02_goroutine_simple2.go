package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// runtime.GOMAXPROCS(1) // Limiting to 1 CPU core

	wg.Add(3)
	go doSomething(10)
	go doSomething(10)
	go doSomething(10)

	wg.Wait()
}

func doSomething(n int) {
	defer wg.Done()

	for i := 0; i < n; i++ {
		fmt.Print(i)
	}
	fmt.Println(" ")
}
