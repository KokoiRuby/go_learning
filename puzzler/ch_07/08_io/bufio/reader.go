package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := strings.NewReader("Hello, せ界!")
	//input := strings.NewReader("Hello, せ界!", 4096)
	reader := bufio.NewReader(input)
	fmt.Printf("size of the underlying buffer in bytes: %v\n", reader.Size())

	data := make([]byte, 5)
	n, err := reader.Read(data)
	if err != nil {
		fmt.Println("Error reading data:", err)
	} else {
		fmt.Printf("Bytes read: %s\n", string(data[:n]))
	}
	fmt.Println()

	nextByte, err := reader.ReadByte()
	if err != nil {
		fmt.Println("Error reading byte:", err)
	} else {
		fmt.Printf("Next byte read: %c\n", nextByte)
	}
	err = reader.UnreadByte()
	if err != nil {
		return
	}
	fmt.Println()

	line, err := reader.ReadBytes(' ')
	if err != nil {
		fmt.Println("Error reading bytes until delimiter:", err)
	} else {
		fmt.Printf("Bytes read until delimiter: %s\n", string(line))
	}
	fmt.Println()

	line, isPrefix, err := reader.ReadLine()
	if err != nil {
		fmt.Println("Error reading line:", err)
	} else {
		fmt.Printf("Line read: %s, IsPrefix: %t\n", string(line), isPrefix)
	}
	fmt.Println()

	input = strings.NewReader("Hello, せ界!")
	reader = bufio.NewReader(input)
	nextRune, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println("Error reading rune:", err)
	} else {
		fmt.Printf("Next rune read: %c\n", nextRune)
	}
	err = reader.UnreadRune()
	if err != nil {
		return
	}
	fmt.Println()

	input = strings.NewReader("Hello, せ界!")
	reader = bufio.NewReader(input)
	line, err = reader.ReadSlice(',')
	if err != nil {
		fmt.Println("Error reading bytes until delimiter:", err)
	} else {
		fmt.Printf("Bytes read until delimiter: %s\n", string(line))
	}

	input = strings.NewReader("Hello, せ界!")
	reader = bufio.NewReader(input)
	lineString, err := reader.ReadString('!')
	if err != nil {
		fmt.Println("Error reading string until delimiter:", err)
	} else {
		fmt.Printf("String read until delimiter: %s\n", lineString)
	}
	fmt.Println()

	//

	input = strings.NewReader("Hello, せ界!")
	reader = bufio.NewReader(input)
	bufferedBytes := reader.Buffered()
	fmt.Printf("Number of bytes that can be read in current buffer: %d\n", bufferedBytes)
	data = make([]byte, 3)
	n, _ = reader.Read(data)
	fmt.Printf("Bytes read: %s\n", string(data[:n]))
	bufferedBytes = reader.Buffered()
	fmt.Printf("Number of bytes that can be read in current buffer: %d\n", bufferedBytes)
	fmt.Println()

	input = strings.NewReader("Hello, World!")
	reader = bufio.NewReader(input)
	discarded, _ := reader.Discard(5)
	fmt.Printf("Discarded %d bytes\n", discarded)
	data = make([]byte, 3)
	n, _ = reader.Read(data)
	fmt.Printf("Bytes read: %s\n", string(data[:n]))
	fmt.Println()

	input = strings.NewReader("Hello, World!")
	reader = bufio.NewReader(input)
	peekedBytes, _ := reader.Peek(5)
	fmt.Printf("Peeked bytes: %s\n", string(peekedBytes))
	fmt.Println()

	input = strings.NewReader("Hello, World!")
	reader = bufio.NewReader(input)
	data = make([]byte, 3)
	n, _ = reader.Read(data)
	fmt.Printf("Bytes read: %s\n", string(data[:n]))
	fmt.Println()

	reader.Reset(input)
	data = make([]byte, 3)
	n, _ = reader.Read(data)
	fmt.Printf("Bytes read: %s\n", string(data[:n]))
	fmt.Println()

	//

	input = strings.NewReader("Hello, World!")
	reader = bufio.NewReader(input)
	bytesWritten, _ := reader.WriteTo(os.Stdout)
	println("\nBytes written to stdout:", bytesWritten)
	fmt.Println()
}
