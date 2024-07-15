package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, Go strings.Reader!"
	reader := strings.NewReader(str)

	_, err := reader.Seek(7, 0)
	if err != nil {
		fmt.Println("Error seeking:", err)
		return
	}

	buf := make([]byte, 5)
	n, err := reader.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Printf("Read %d bytes after seeking: %s\n", n, string(buf[:n]))
}
