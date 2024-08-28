package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./sample.txt")
	if err != nil {
		fmt.Println("Error opening file.go:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

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
				fmt.Println("Reached end of file.go.")
				break
			} else {
				fmt.Println("Error reading file.go:", err)
				break
			}
		}
		fmt.Print(string(buf[:n]))
	}
}
