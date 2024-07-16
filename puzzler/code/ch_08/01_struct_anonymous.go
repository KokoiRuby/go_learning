package main

import "fmt"

type D struct {
	string
	int
}

func main() {
	p := struct {
		Name  string
		Age   int
		Email string
	}{
		Name:  "Alice",
		Age:   25,
		Email: "alice@example.com",
	}

	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Email:", p.Email)

	d := D{"Alice", 25}
	fmt.Println(d.string)
	fmt.Println(d.int)

}
