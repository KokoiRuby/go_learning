package main

import "fmt"

func modify(arr *[3]int) {
	(*arr)[0] = 4
}

func main() {
	arr := [3]int{1, 2, 3}
	modify(&arr)
	fmt.Println(arr)
}
