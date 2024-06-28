package main

import "fmt"

func main() {
	a := func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(a)
}
