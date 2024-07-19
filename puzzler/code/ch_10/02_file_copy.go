package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	// 1. open src
	src, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer func(src *os.File) {
		err := src.Close()
		if err != nil {

		}
	}(src)

	// 2. create dst
	dst, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {

		}
	}(dst)

	// 3. create reader & writer
	reader := bufio.NewReader(src)
	writer := bufio.NewWriter(dst)

	return io.Copy(writer, reader)

}

func main() {
	src := "./sample.txt"
	dst := "./sample_copy.txt"
	_, err := CopyFile(dst, src)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Copy successful!")
}
