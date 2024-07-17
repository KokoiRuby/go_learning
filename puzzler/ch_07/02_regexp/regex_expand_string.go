package main

import (
    "fmt"
    "regexp"
)

func main() {
    src := "The email is example@example.com"
    re := regexp.MustCompile(`(\w+)@(\w+)\.(\w+)`)
    template := "domain=${2}.${3}"
    dst := ""
    result := re.ExpandString([]byte(dst), template, src, re.FindStringSubmatchIndex(src))
    fmt.Println(string(result)) // domain=example.com
}
