package main

import "fmt"

func main() {
	var name string
	var age int
	var gender bool
	fmt.Println("Please input name, age, gender: ")
	fmt.Scanln(&name)
	fmt.Scanln(&age)
	fmt.Scanln(&gender)
	fmt.Printf("name is: %v\n", name)
	fmt.Printf("age is: %v\n", age)
	fmt.Printf("gender is: %v\n", gender)
}
