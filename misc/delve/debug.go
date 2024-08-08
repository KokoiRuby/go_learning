package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")

	a := 10
	b := 20
	sum := add(a, b)

	fmt.Printf("Sum of %d and %d is %d\n", a, b, sum)

	fmt.Println("End")
}

func add(x, y int) int {
	return x + y
}
