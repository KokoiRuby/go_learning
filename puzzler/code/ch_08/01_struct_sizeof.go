package main

import (
	"fmt"
	"unsafe"
)

type MyStruct struct {
	Field1 int
	Field2 string
	Field3 bool
}

func main() {
	instance := MyStruct{
		Field1: 10,
		Field2: "Hello",
		Field3: true,
	}

	size := unsafe.Sizeof(instance)
	fmt.Printf("MyStruct instance size is: %d bytes\n", size)
}
