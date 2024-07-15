package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	str := "Hello, Go strings.Reader!"
	reader := strings.NewReader(str)

	buf := make([]byte, 4)
	n, err := reader.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Printf("Read %d bytes: %s\n", n, string(buf[:n]))

	readingIndex := reader.Size() - int64(reader.Len())
	fmt.Printf("Reading index: %d", readingIndex)

}
