package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Program name:", os.Args[0])

	fmt.Println("Arguments:")
	for i, arg := range os.Args[1:] {
		fmt.Printf("Arg %d: %s\n", i+1, arg)
	}
}
