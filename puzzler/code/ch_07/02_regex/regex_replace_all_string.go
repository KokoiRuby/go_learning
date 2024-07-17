package main

import (
    "fmt"
    "regexp"
)

func main() {
    re := regexp.MustCompile(`apples`)
    str := "There are 123 apples"
    result := re.ReplaceAllString(str, "oranges")
    fmt.Println(result) // There are 123 oranges
}
