package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 10
	num := 11
	intChan <- num

	fmt.Println(len(intChan))
	fmt.Println(cap(intChan))
	intChan <- 12

	// drop
	<-intChan
	<-intChan

	// assign
	n := <-intChan
	fmt.Println(n)
}
