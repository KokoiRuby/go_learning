package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var buf bytes.Buffer
	buf.Grow(64)
	fmt.Printf("Buffer capacity: %d\n", buf.Cap())

	data := []byte("Hello, ")
	n, err := buf.Write(data)
	if err != nil {
		fmt.Println("Error writing bytes to buffer:", err)
	}
	fmt.Printf("bytes written: %d\n", n)
	fmt.Printf("Number of bytes in buffer: %d\n", buf.Len())
	fmt.Println(buf.String())
	fmt.Println()

	err = buf.WriteByte('W')
	if err != nil {
		fmt.Println("Error writing byte to buffer:", err)
	}
	fmt.Printf("Number of bytes in buffer: %d\n", buf.Len())
	fmt.Println(buf.String())
	fmt.Println()

	n, err = buf.WriteRune('せ')
	if err != nil {
		fmt.Println("Error writing rune to buffer:", err)
	}
	fmt.Printf("bytes written: %d\n", n)
	fmt.Printf("Number of bytes in buffer: %d\n", buf.Len())
	fmt.Println(buf.String())
	fmt.Println()

	n, err = buf.WriteString("!")
	if err != nil {
		fmt.Println("Error writing string to buffer:", err)
	}
	fmt.Printf("bytes written: %d\n", n)
	fmt.Printf("Number of bytes in buffer: %d\n", buf.Len())
	fmt.Println(buf.String())
	fmt.Println()

	nWritten, err := buf.WriteTo(os.Stdout)
	if err != nil {
		fmt.Println("Error writing buffer to output:", err)
	} else {
		fmt.Printf("\nbytes written: %d\n", nWritten)
	}
	// since buf writes to buff
	fmt.Printf("Number of bytes in buffer: %d\n", buf.Len())
	fmt.Println(buf.String())
	fmt.Println()

	//

	buf.Write([]byte("Hello, Wせ!"))

	available := buf.Available()
	fmt.Printf("Available bytes in buffer: %d\n", available)

	emptyBuffer := buf.AvailableBuffer()
	fmt.Printf("Empty buffer with capacity %d created\n", cap(emptyBuffer))

	nextBytes := buf.Next(5)
	fmt.Printf("Next 5 bytes from buffer: %s\n", string(nextBytes))
	fmt.Println()

	//

	buf.Reset()
	buf.Write([]byte("Hello, Wせ!"))
	readBytes := make([]byte, 5)
	n, err = buf.Read(readBytes)
	if err != nil {
		fmt.Println("Error reading bytes from buffer:", err)
	} else {
		fmt.Printf("Bytes read: %s\n", string(readBytes[:n]))
	}
	fmt.Println()

	buf.Reset()
	buf.Write([]byte("Hello, Wせ!"))
	nextByte, err := buf.ReadByte()
	if err != nil {
		fmt.Println("Error reading byte from buffer:", err)
	} else {
		fmt.Printf("Next byte read: %c\n", nextByte)
	}
	err = buf.UnreadByte()
	if err != nil {
		return
	}
	fmt.Printf(buf.String())
	fmt.Println()

	buf.Reset()
	buf.Write([]byte("Hello, Wせ!"))
	line, err := buf.ReadBytes(' ')
	if err != nil {
		fmt.Println("Error reading bytes until delimiter from buffer:", err)
	} else {
		fmt.Printf("Bytes read until delimiter: %s\n", string(line))
	}
	fmt.Println()

	buf.Reset()
	buf.Write([]byte("Hello, Wせ!"))
	readerData := []byte(" This is additional data.")
	num, err := buf.ReadFrom(bytes.NewReader(readerData))
	if err != nil {
		fmt.Println("Error reading from reader:", err)
	} else {
		fmt.Printf("Bytes read from reader: %d\n", num)
	}
	fmt.Println()

	buf.Reset()
	buf.Write([]byte("Hello, Wせ!"))
	nextRune, _, err := buf.ReadRune()
	if err != nil {
		fmt.Println("Error reading rune from buffer:", err)
	} else {
		fmt.Printf("Next rune read: %c\n", nextRune)
	}
	err = buf.UnreadRune()
	if err != nil {
		return
	}
	fmt.Printf(buf.String())
	fmt.Println()

	buf.Reset()
	buf.Write([]byte("Hello, Wせ!"))
	lineString, err := buf.ReadString('!')
	if err != nil {
		fmt.Println("Error reading string until delimiter from buffer:", err)
	} else {
		fmt.Printf("String read until delimiter: %s\n", lineString)
	}
	fmt.Println()
}
