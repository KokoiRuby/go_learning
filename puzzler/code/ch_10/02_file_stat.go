package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	az    int
	num   int
	space int
	other int
}

func main() {
	path := "./sample.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var count CharCount
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range line {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.az++
			case v == ' ' || v == '\t':
				count.space++
			case v >= '0' && v <= '9':
				count.num++
			default:
				count.other++
			}
		}
	}
	fmt.Println(count)
}
