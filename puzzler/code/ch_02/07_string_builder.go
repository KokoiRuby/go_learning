package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder

	// Initial writes
	b.WriteString("Hello")
	b.WriteString(" ")
	b.WriteString("World")

	// Print the current content and length
	fmt.Println("Current content:", b.String()) // Output: Hello World
	fmt.Println("Current length:", b.Len())     // Output: 11

	// Write a longer string to trigger buffer expansion
	longString := " and welcome to the Go programming language. This is a demonstration of strings.Builder in Go."
	b.WriteString(longString)

	// Print the current content and length
	fmt.Println("Current content:", b.String()) // Output: Hello World and welcome to the Go programming language. This is a demonstration of strings.Builder in Go.
	fmt.Println("Current length:", b.Len())     // Output: 113

	// Write more data to see the expansion in action
	moreString := " The builder will expand its internal buffer as needed."
	b.WriteString(moreString)

	// Print the final content and length
	fmt.Println("Final content:", b.String()) // Output: Hello World and welcome to the Go programming language. This is a demonstration of strings.Builder in Go. The builder will expand its internal buffer as needed.
	fmt.Println("Final length:", b.Len())     // Output: 164

	// Demonstrate the Grow method
	var b2 strings.Builder
	b2.Grow(64) // Pre-allocate space for 64 bytes
	b2.WriteString("Pre-allocated buffer space with Grow method.")
	fmt.Println("Builder with pre-allocated space:", b2.String())
	fmt.Println("Length of pre-allocated builder:", b2.Len())

	// Reset the builder
	b.Reset()

	// Print the current content and length after reset
	fmt.Println("Content after reset:", b.String()) // Output: (empty string)
	fmt.Println("Length after reset:", b.Len())     // Output: 0

	// Write new data after reset
	b.WriteString("New content after reset")
	fmt.Println("New content:", b.String()) // Output: New content after reset
	fmt.Println("New length:", b.Len())     // Output: 22
}
