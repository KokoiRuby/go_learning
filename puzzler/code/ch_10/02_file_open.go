package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./sample.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
}
