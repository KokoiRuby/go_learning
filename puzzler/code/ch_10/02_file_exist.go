package main

import (
	"fmt"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}

func main() {
	path := "./sample.txt"
	_, err := PathExists(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Path exists.")
}
