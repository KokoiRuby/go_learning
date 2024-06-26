package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./sample.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// buffer
	buf := make([]byte, 1024)

	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == os.ErrClosed {
				fmt.Println("File closed.")
				break
			} else if err == io.EOF {
				if n > 0 {
					fmt.Print(string(buf[:n]))
				}
				fmt.Println("Reached end of file.")
				break
			} else {
				fmt.Println("Error reading file:", err)
				break
			}
		}
		fmt.Print(string(buf[:n]))
	}
}
