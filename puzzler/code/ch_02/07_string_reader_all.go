package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	str := "Hello, Go strings.Reader!"
	reader := strings.NewReader(str)

	var result []byte
	buf := make([]byte, 4)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		result = append(result, buf[:n]...)
	}

	fmt.Printf("Read all content: %s\n", string(result))
}
