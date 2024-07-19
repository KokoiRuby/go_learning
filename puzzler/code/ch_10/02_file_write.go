package main

import (
	"bufio"
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

	// write
	str := "Write something...\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		_, err := writer.WriteString(str)
		if err != nil {
			return
		}
	}
	err = writer.Flush()
	if err != nil {
		return
	}
}
