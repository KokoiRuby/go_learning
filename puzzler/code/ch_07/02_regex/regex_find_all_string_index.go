package main

import (
    "fmt"
    "regexp"
)

func main() {
    re := regexp.MustCompile(`\d+`)
    str := "There are 123 apples and 456 bananas"
    matches := re.FindAllStringIndex(str, -1)
    fmt.Println(matches) // [[10 13] [23 26]]
}
