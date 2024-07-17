package main

import (
    "fmt"
    "regexp"
)

func main() {
    src := []byte("The email is example@example.com")
    re := regexp.MustCompile(`(\w+)@(\w+)\.(\w+)`)
    template := []byte("domain=${2}.${3}")
    dst := []byte{}
    result := re.Expand(dst, template, src, re.FindSubmatchIndex(src))
    fmt.Println(string(result)) // domain=example.com
}
