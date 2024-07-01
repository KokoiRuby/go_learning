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
	defer file.Close()
}
