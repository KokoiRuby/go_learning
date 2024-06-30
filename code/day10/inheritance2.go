package main

import "fmt"

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
}

func main() {
	var b B
	b.A.Name = "Jack"
	b.Name = "Tom"
	b.A.age = 18
	b.age = 18
	b.A.SayOk()
	b.SayOk()
	b.A.hello()
	b.hello()

}
