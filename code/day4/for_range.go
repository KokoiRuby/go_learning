package main

import "fmt"

func main() {
	strings := []string{"google", "amazon"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 4}
	for i, n := range numbers {
		fmt.Println(i, n)
	}
}
