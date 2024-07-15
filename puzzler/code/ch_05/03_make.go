package main

import "fmt"

func main() {
	// slice
	s := make([]int, 5)
	fmt.Println(s)

	// map
	m := make(map[string]string)
	m["key"] = "value"
	fmt.Println(m)

	// channel
	c := make(chan int, 1)
	c <- 1
	fmt.Println(<-c)
}
