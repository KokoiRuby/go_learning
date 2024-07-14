package main

import "fmt"

func main() {
	str1 := "Ghost of 対馬"
	runes := []rune(str1)

	for i, r := range runes {
		fmt.Printf("Character at index %d is %c\n", i, r)
	}

	rune := '対'
	str2 := string(rune)
	fmt.Println("String repr is: ", str2)
}
