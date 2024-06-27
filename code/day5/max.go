package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Input a number: ")
	fmt.Scanln(&a)
	fmt.Println("Input another number: ")
	fmt.Scanln(&b)

	fmt.Println(max(a, b))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
