package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x int = 42
	ptr1 := &x
	uintptrPtr := uintptr(unsafe.Pointer(ptr1)) // unsafe.Pointer → uintptr
	fmt.Printf("Address of x: %v\n", uintptrPtr)

	var y int = 42
	ptr2 := unsafe.Pointer(&y)
	fmt.Printf("%v\n", ptr2)
	val1 := *(*int)(ptr2) // unsafe.Pointer → *int then take val
	fmt.Printf("Value of x: %v\n", val1)
}
