package main

import "fmt"

type Boss struct {
	name string
	age  int
}

func (p *Boss) GetName() string {
	return p.name
}

func (p *Boss) SetName(name string) {
	p.name = name
}

func (p *Boss) GetAge() int {
	return p.age
}

func (p *Boss) SetAge(age int) {
	p.age = age
}

func main() {
	Boss := Boss{
		name: "Alice",
		age:  25,
	}

	fmt.Println("Name:", Boss.GetName())
	fmt.Println("Age:", Boss.GetAge())

	Boss.SetName("Bob")
	Boss.SetAge(30)

	fmt.Println("Name:", Boss.GetName())
	fmt.Println("Age:", Boss.GetAge())
}
