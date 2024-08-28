package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	str := "こんにちは World"
	strReader := strings.NewReader(str)
	fmt.Printf("avail number of bytes for reading: %v\n", strReader.Size())

	var result []byte
	// buffer size
	buf := make([]byte, 4)
	// read per buffer size until EOF
	for {
		n, err := strReader.Read(buf)
		fmt.Printf("number of bytes read: %v\n", n)
		fmt.Printf("total number of bytes read: %v\n", strReader.Size()-int64(strReader.Len()))
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		fmt.Println(string(buf[:n]))
		result = append(result, buf[:n]...)
	}

	fmt.Printf("Read all content: %s\n", string(result))
	fmt.Println()

	//

	strReader.Reset(str)
	buf = make([]byte, 4)
	n, err := strReader.ReadAt(buf, 7)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Printf("Read %d bytes from idx 7: %s\n", n, string(buf[:n]))
	fmt.Println()

	//

	strReader.Reset(str)
	for {
		b, err := strReader.ReadByte()
		fmt.Printf("byte read: %v\n", b)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
	}
	fmt.Println()

	//

	strReader.Reset(str)
	for {
		r, size, err := strReader.ReadRune()
		fmt.Printf("rune read: %v, size in byte: %v\n", r, size)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
	}
	fmt.Println()

	//

	strReader.Reset(str)
	b, err := strReader.ReadByte()
	if err != nil {
		fmt.Println("Error reading:", err)
	}
	fmt.Printf("byte read: %v\n", b)
	err = strReader.UnreadByte()
	if err != nil {
		fmt.Println("Error while unread rune:", err)
	}
	b, err = strReader.ReadByte()
	if err != nil {
		fmt.Println("Error reading:", err)
	}
	fmt.Printf("byte read after byte unread: %v\n", b)
	fmt.Println()

	//

	strReader.Reset(str)
	r, size, err := strReader.ReadRune()
	if err != nil {
		fmt.Println("Error reading:", err)
	}
	fmt.Printf("rune read: %v, size in byte: %v\n", r, size)
	err = strReader.UnreadRune()
	if err != nil {
		fmt.Println("Error while unread rune:", err)
	}
	r, size, err = strReader.ReadRune()
	if err != nil {
		fmt.Println("Error reading:", err)
	}
	fmt.Printf("rune read: %v, size in byte: %v\n", r, size)
	fmt.Println()

	//

	strReader.Reset(str)
	num, err := strReader.WriteTo(os.Stdout)
	if err != nil {
		fmt.Println("Error writing to output:", err)
	} else {
		fmt.Printf("\nTotal bytes written: %d\n", num)
	}
	fmt.Println()

	//

	strReader.Reset(str)
	offset, err := strReader.Seek(7, 0)
	//offset, err := strReader.Seek(7, 1)
	//offset, err := strReader.Seek(-7, 2)
	if err != nil {
		fmt.Println("Error while seeking:", err)
	} else {
		fmt.Printf("New offset after seek: %d\n", offset)

		b, err := strReader.ReadByte()
		if err != nil {
			fmt.Println("Error reading:", err)
		}
		fmt.Printf("byte read: %v\n", b)
	}

}
