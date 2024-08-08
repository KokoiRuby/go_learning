package main

import "fmt"

// Filter iter int arr & call fn to filter then append to a new int arr
// fn returns bool as cond
func Filter(arr []int, fn func(n int) bool) []int {
	var newArray []int
	for _, it := range arr {
		if fn(it) {
			newArray = append(newArray, it)
		}
	}
	return newArray
}

func main() {
	var s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	out := Filter(s, func(n int) bool {
		return n%2 == 1
	})
	fmt.Printf("%v\n", out)
	out = Filter(s, func(n int) bool {
		return n > 5
	})
	fmt.Printf("%v\n", out)
}