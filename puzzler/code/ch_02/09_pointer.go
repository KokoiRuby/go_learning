package main

import (
	"fmt"
)

func main() {
	var a int = 20
	var ip *int

	ip = &a

	fmt.Printf("addr of a is: %x\n", &a)
	fmt.Printf("addr stored in ptr is: %x\n", ip)
	fmt.Printf("value ptr points to is: %x\n", *ip)

	var ptr *int
	fmt.Printf("value of ptr : %x\n", ptr)
	if ptr == nil {
		fmt.Printf("ptr is nil.")
	}
}
