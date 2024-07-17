package main

import (
    "fmt"
    "regexp"
)

func main() {
    re := regexp.MustCompile(`\d+`)
    str := "There are 123 apples and 456 bananas"
    matches := re.FindAllString(str, -1) // -1 means all
    fmt.Println(matches) // [123 456]
	matches = re.FindAllString(str, 1)
	fmt.Println(matches) // [123]
	matches = re.FindAllString(str, 2)
	fmt.Println(matches) // [123 456]
}
