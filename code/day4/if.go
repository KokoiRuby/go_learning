package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("Input a: ")
	fmt.Scan(&a)
	fmt.Println("Input b: ")
	fmt.Scan(&b)

	// determine if a+b can be divided by both 3 and 5
	var sum int = a + b
	if sum%3 == 0 && sum%5 == 0 {
		fmt.Println("sum can be divided by both 3 and 5.")
	} else {
		fmt.Println("sum cannot be divided by both 3 and 5.")
	}
}
