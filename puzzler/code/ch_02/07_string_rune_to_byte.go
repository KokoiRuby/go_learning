package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	r1 := '対'
	var byteSlice []byte = make([]byte, utf8.UTFMax)
	n := utf8.EncodeRune(byteSlice, r1)
	fmt.Printf("Byte repr is: %v", byteSlice[:n])

	b := []byte{229, 175, 190}
	r2, size := utf8.DecodeRune(b)
	fmt.Printf("Rune repr is: %c\n Rune size is: %v", r2, size)
}
