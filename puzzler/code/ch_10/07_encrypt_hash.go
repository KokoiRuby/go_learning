package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	data := []byte("hello, world")
	hash := sha256.Sum256(data)
	fmt.Printf("SHA256 hash: %x\n", hash)
}
