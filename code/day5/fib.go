package main

import "fmt"

func fib(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}

}

func main() {
	var n int
	fmt.Println("Input a number: ")
	fmt.Scanln(&n)
	result := fib(n)
	fmt.Println(result)

}
