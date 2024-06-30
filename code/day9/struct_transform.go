package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Salary float64
}

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	p := Person{
		Name: "john",
		Age:  20,
	}
	// Person → Employee
	var e Employee
	e = Employee(p)
	fmt.Println(e, p)

}
