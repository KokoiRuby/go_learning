package main

import "fmt"

func cal(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return sum, sub
}

func main() {
	a, b := cal(3, 2)
	fmt.Println(a, b)
}
