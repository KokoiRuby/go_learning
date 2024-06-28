package main

import "fmt"

func getSum(a, b int) int {
	return a + b
}

func main() {
	x := getSum
	fmt.Printf("type of var a is: %T", x)
}
