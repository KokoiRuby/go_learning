package main

import "fmt"

func main() {
	str := "42 John"

	var num int
	var name string

	n, err := fmt.Sscanf(str, "%d %s", &num, &name)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Number:", num)
	fmt.Println("Name:", name)
	fmt.Println("Items scanned:", n)
}
