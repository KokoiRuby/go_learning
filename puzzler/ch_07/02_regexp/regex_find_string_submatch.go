package main

import (
    "fmt"
    "regexp"
)

func main() {
    re := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
    str := "Date: 2023-07-15"
    match := re.FindStringSubmatch(str)
    fmt.Println(match) // [2023-07-15 2023 07 15]
}
