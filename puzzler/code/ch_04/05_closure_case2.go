package main

import "fmt"

func foo7(x int) []func() {
	var fs []func() // func slice
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fs = append(fs, func() {
			fmt.Printf("foo7 val = %d\n", x+val)
		})
	}
	return fs
}

func main() {
	f7s := foo7(11)
	for _, f7 := range f7s {
		f7() // 12 13 14 16
	}
}
