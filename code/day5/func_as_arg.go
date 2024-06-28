package main

import "fmt"

func getSum(a, b int) int {
	return a + b
}

func myFunc(funvar func(int, int) int, a, b int) int {
	return funvar(a, b)
}

func main() {
	b := myFunc(getSum, 1, 2)
	fmt.Println(b)
}
