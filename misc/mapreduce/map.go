package main

import (
	"fmt"
	"strconv"
	"strings"
)

// MapStrToStr iter string arr & call fn then append to a new string arr
func MapStrToStr(arr []string, fn func(s string) string) []string {
	var newArray []string
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

// MapStrToInt iter a string arr & call fn then append to a new int arr
func MapStrToInt(arr []string, fn func(s string) int) []int {
	var newArray []int
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func main() {
	var list = []string{"a", "b", "c"}

	x := MapStrToStr(list, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Printf("%v\n", x)
	//["A", "B", "C"]

	list = []string{"1", "2", "3"}
	y := MapStrToInt(list, func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	})
	fmt.Printf("%v\n", y)
	//[1, 2, 3]
}
