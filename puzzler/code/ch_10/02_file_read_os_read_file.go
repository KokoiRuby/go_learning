package main

import (
	"fmt"
	"os"
)

func main() {
	byteString, err := os.ReadFile("./sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteString))
}
