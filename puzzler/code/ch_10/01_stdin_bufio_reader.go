package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// read from stdin
	reader := bufio.NewReader(os.Stdin)

	for {
		// read per line
		line, err := reader.ReadString('\n')

		// err handling
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, "Error reading input:", err)
			if err == io.EOF {
				return
			}
			break
		}

		fmt.Println("Input:", line)

		if line == "\n" {
			break
		}
	}
}
