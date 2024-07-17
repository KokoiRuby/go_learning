package main

import (
    "fmt"
    "regexp"
)

func main() {
    re := regexp.MustCompile(`\d+`)
    str := "There are 123 apples"
    match := re.FindString(str)
    fmt.Println(match) // 123
}
