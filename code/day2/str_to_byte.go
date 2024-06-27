package main

import "fmt"

func main() {
	str1 := "hello"
	byteSlice := []byte(str1)

	for i, b := range byteSlice {
		fmt.Printf("Character at index %d is %v\n", i, b)
	}

	byte := 98
	str2 := string(byte)
	fmt.Println("String repr is: ", str2)

}
