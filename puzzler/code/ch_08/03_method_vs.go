package main

import "fmt"

type Animal struct {
	Name string
}

func test01(a Animal) {
	fmt.Println(a.Name)
}

func test02(a *Animal) {
	fmt.Println(a.Name)
}

func (a Animal) test03() {
	a.Name = "Jack"
	fmt.Println(a.Name)
}

func (a *Animal) test04() {
	a.Name = "Jack"
	fmt.Println(a.Name)
}

func main() {
	a := Animal{"Tom"}
	test01(a)  // Tom
	test02(&a) // Tom

	a.test03()          // Jack
	fmt.Println(a.Name) // Tom

	(&a).test03()       // Jack
	fmt.Println(a.Name) // Tom

	a.test04()          // Jack
	fmt.Println(a.Name) // Jack

	(&a).test04()       // Jack
	fmt.Println(a.Name) // Jack
}
