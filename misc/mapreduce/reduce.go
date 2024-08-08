package main

import "fmt"

// Reduce iter string arr & call fn to then sum
func Reduce(arr []string, fn func(s string) int) int {
	sum := 0
	for _, it := range arr {
		sum += fn(it)
	}
	return sum
}

func main() {
	var list = []string{"a", "bb", "ccc"}
	x := Reduce(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("%v\n", x)
	// 6
}
