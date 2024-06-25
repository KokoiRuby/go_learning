package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n = 1000000
	fmt.Printf("size of type of n is: %d", unsafe.Sizeof(n))
}
