package main

import "fmt"

func sum(numbers ...int) int {
	s := 0
	for i := 0; i < len(numbers); i++ {
		s += numbers[i]
	}
	return s
}

func main() {
	a := sum(1, 2, 3, 4, 5)
	fmt.Println(a)
}
