package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str1 string = "true"
	var b bool
	b, _ = strconv.ParseBool(str1)
	fmt.Printf("type is: %T, value is: %v\n", b, b)

	var str2 string = "123"
	var n int64
	n, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("type is: %T, value is: %v\n", n, n)
}
