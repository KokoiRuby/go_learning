package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	//writer := bufio.NewWriterSize(os.Stdout, 4096)
	fmt.Printf("size of the underlying buffer in bytes: %v\n", writer.Size())

	bytes := []byte{'H', 'e', 'l', 'l', 'o'}
	n, _ := writer.Write(bytes)
	fmt.Println("Bytes written:", n)
	availableBytes := writer.Available()
	fmt.Printf("Available bytes in buffer: %d\n", availableBytes)
	fmt.Println()

	_ = writer.WriteByte(' ')
	availableBytes = writer.Available()
	fmt.Printf("Available bytes in buffer: %d\n", availableBytes)
	fmt.Println()

	size, _ := writer.WriteRune('世')
	fmt.Println("Bytes written for rune:", size)
	availableBytes = writer.Available()
	fmt.Printf("Available bytes in buffer: %d\n", availableBytes)
	fmt.Println()

	n, _ = writer.WriteString("界!\n")
	fmt.Println("Bytes written for string:", n)
	availableBytes = writer.Available()
	fmt.Printf("Available bytes in buffer: %d\n", availableBytes)
	emptyBuffer := writer.AvailableBuffer()
	fmt.Printf("Empty buffer capacity: %d\n", cap(emptyBuffer))
	fmt.Println()

	_ = writer.Flush()
	fmt.Println()

	//

	bytes = []byte{'H', 'e', 'l', 'l', 'o'}
	n, _ = writer.Write(bytes)
	fmt.Println("Bytes written:", n)
	availableBytes = writer.Available()
	fmt.Printf("Available bytes in buffer: %d\n", availableBytes)
	writer.Reset(os.Stdout)
	availableBytes = writer.Available()
	fmt.Printf("Available bytes in buffer after reset: %d\n", availableBytes)
	fmt.Println()

	strReader := strings.NewReader("Example text to be read from.")
	bytesRead, _ := writer.ReadFrom(strReader)
	fmt.Printf("Bytes read and written: %v", bytesRead)

}
