package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, Go strings.Reader!"
	reader := strings.NewReader(str)

	buf := make([]byte, 5)
	n, err := reader.ReadAt(buf, 7)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Printf("Read %d bytes from position 7: %s\n", n, string(buf[:n]))
}
