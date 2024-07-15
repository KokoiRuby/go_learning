package main

import "fmt"

func main() {
	sum := func(a, b int) int {
		return a + b
	}
	fmt.Println(sum(1, 2))
	fmt.Println(sum(3, 4))
}
