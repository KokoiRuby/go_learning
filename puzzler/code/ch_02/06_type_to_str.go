package main

import (
	"fmt"
	"strconv"
)

func main() {
	var stockcode = 123
	var enddate = "2024-06-27"
	var url = "Code=%d&endDate=%s"
	// format and return a string without printing it
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
	// fmt.Printf(target_url, stockcode, enddate)

	var num1 int = 99
	var num2 float64 = 23.4325
	var b bool = true
	var myChar byte = 'h'
	var str string

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("type is: %T, value is: %v\n", str, str)
	str = fmt.Sprintf("%f", num2)
	fmt.Printf("type is: %T, value is: %v\n", str, str)
	str = fmt.Sprintf("%t", b)
	fmt.Printf("type is: %T, value is: %v\n", str, str)
	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("type is: %T, value is: %v\n", str, str)

	var num3 int = 456
	str = strconv.Itoa(num3)
	fmt.Printf("type is: %T, value is: %v\n", str, str)

	var num4 int64 = 123
	// str = strconv.FormatInt(num4, 10)
	str = strconv.FormatInt(num4, 2)
	fmt.Printf("type is: %T, value is: %v\n", str, str)
}
