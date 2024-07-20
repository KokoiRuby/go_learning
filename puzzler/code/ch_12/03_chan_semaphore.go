package main

import (
	"fmt"
	"time"
)

func worker(id int, sem chan struct{}) {
	fmt.Printf("Worker %d acquiring resource\n", id)
	<-sem // recv from chan, semaphore--

	fmt.Printf("Worker %d accessing resource\n", id)
	time.Sleep(time.Second)

	fmt.Printf("Worker %d releasing resource\n", id)
	sem <- struct{}{} // send to chan, semaphore++
}

func main() {
	const maxConcurrency = 3
	sem := make(chan struct{}, maxConcurrency) // buffered chan to control max concurrency

	// fill in
	sem <- struct{}{}
	sem <- struct{}{}
	sem <- struct{}{}

	for i := 1; i <= 5; i++ {
		go worker(i, sem)
	}

	time.Sleep(3 * time.Second) // wait until all goroutines

	close(sem) // close chan
	time.Sleep(time.Second)
}
