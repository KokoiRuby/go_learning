package main

import (
    "fmt"
    "regexp"
)

func main() {
    re := regexp.MustCompile(`(?P<Year>\d{4})-(?P<Month>\d{2})-(?P<Day>\d{2})`)
    str := "Date: 2023-07-15"
    match := re.FindStringSubmatch(str)
    
    // iter named capture group
    for i, name := range re.SubexpNames() {
        if i != 0 && name != "" {
            fmt.Printf("%s: %s\n", name, match[i])
        }
    }
    // Year: 2023
    // Month: 07
    // Day: 15
}
