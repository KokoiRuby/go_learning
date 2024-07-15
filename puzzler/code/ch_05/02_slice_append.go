package main

import "fmt"

func main() {
	sl := []int{1, 2, 3}
	newSl := append(sl, 4, 5)
	fmt.Println(sl)
	fmt.Println(newSl)
	fmt.Println(cap(newSl))
	fmt.Println(len(newSl))
}
