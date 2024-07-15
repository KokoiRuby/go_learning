package main

import "fmt"

func main() {
	var num *int
	num = new(int)
	fmt.Printf("%T\n", num)
	fmt.Printf("%v\n", num)
	fmt.Printf("%v\n", &num)
	fmt.Printf("%v\n", *num)
}
