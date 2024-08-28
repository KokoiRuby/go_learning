package main

import (
	"fmt"
	"strings"
)

func main() {
	var strBuilder strings.Builder

	strBuilder.WriteString("Hello")
	strBuilder.WriteString(" ")
	strBuilder.WriteString("world")

	fmt.Println(strBuilder.String())
	fmt.Println(strBuilder.Len())
	fmt.Println(strBuilder.Cap())
	strBuilder.Reset()
	fmt.Println(strBuilder.String())

	strBuilder.WriteByte('h')
	strBuilder.WriteByte('e')
	strBuilder.WriteByte('l')
	strBuilder.WriteByte('l')
	strBuilder.WriteByte('o')
	fmt.Println(strBuilder.String())
	fmt.Println(strBuilder.Len())
	fmt.Println(strBuilder.Cap())
	strBuilder.Reset()
	fmt.Println(strBuilder.String())

	strBuilder.WriteRune('こ')
	strBuilder.WriteRune('ん')
	strBuilder.WriteRune('に')
	strBuilder.WriteRune('ち')
	strBuilder.WriteRune('は')
	fmt.Println(strBuilder.String())
	fmt.Println(strBuilder.Len())
	fmt.Println(strBuilder.Cap())
	strBuilder.Reset()
	fmt.Println(strBuilder.String())

	var content = []byte{'h', 'e', 'l', 'l', 'o'}
	cnt, _ := strBuilder.Write(content)
	fmt.Printf("number of bytes written: %d\n", cnt)
	fmt.Println(strBuilder.String())
	fmt.Println(strBuilder.Len())
	fmt.Println(strBuilder.Cap())
	strBuilder.Reset()
	fmt.Println(strBuilder.String())

	strBuilder.Grow(23)
	fmt.Println(strBuilder.Len())
	fmt.Println(strBuilder.Cap())

}
