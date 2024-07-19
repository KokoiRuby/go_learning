package main

import "fmt"

func main() {
	var num int
	var name string

	fmt.Print("Enter a number: ")
	_, err := fmt.Scanln(&num)
	if err != nil {
		return
	}

	fmt.Print("Enter your name: ")
	_, err = fmt.Scanln(&name)
	if err != nil {
		return
	}

	fmt.Println("Number:", num)
	fmt.Println("Name:", name)
}
