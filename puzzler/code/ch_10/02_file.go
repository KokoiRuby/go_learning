package main

import (
	"fmt"
	"os"
)

func main() {
	// open
	file, err := os.Open("./sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("file.go is %v", file)

	// close
	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}
}
