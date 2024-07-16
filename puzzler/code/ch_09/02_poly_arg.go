package main

import "fmt"

type A interface {
	String() string
	GetName() string
	SetName(string)
}

type A1 struct {
	name string
	age  int64
}

func (a *A1) GetName() string {
	return a.name
}

func (a *A1) SetName(name string) {
	a.name = name
}

func (a *A1) String() string {
	return fmt.Sprintf("name=%v, age=%v", a.name, a.age)
}

func main() {
	a1 := &A1{
		name: "Tom",
		age:  20,
	}

	var a A = a1
	fmt.Println(a)
	fmt.Println(a.GetName())
	a.SetName("Jack")
	fmt.Println(a.GetName())
	fmt.Println(a)

}
