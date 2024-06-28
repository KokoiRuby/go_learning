package main

import "fmt"

var (
	MyFunc = func(a, b int) int {
		return a + b
	}
)

func main() {
	a := MyFunc(1, 2)
	fmt.Println(a)
}
