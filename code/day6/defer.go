package main

import "fmt"

func sum(a, b int) int {
	// push
	defer fmt.Println(a) // 1
	defer fmt.Println(b) // 2
	s := a + b
	fmt.Println(s)
	// pop
	return s
}

func main() {
	fmt.Println(sum(1, 2))
}
