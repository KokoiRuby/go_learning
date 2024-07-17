package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var value atomic.Value

	value.Store(42)
	fmt.Println("Stored value:", value.Load())

	value.Store("Hello, World!")
	fmt.Println("Stored value:", value.Load())
}
