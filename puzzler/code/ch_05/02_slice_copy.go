package main

import "fmt"

func main() {
	var sl []int = []int{1, 2, 3, 4, 5}
	var newsl = make([]int, 2)
	copy(newsl, sl)

	fmt.Println(sl)
	fmt.Println(newsl)

	newsl[1] = 11
	fmt.Println(sl)
	fmt.Println(newsl)
}
