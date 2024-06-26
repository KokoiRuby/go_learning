package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println(line)
				break
			}
			fmt.Println(err)
			break
		}
		fmt.Println(line)
	}

}
