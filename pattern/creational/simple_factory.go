package main

import "fmt"

/*
Abstraction Layer
*/

type Fruit interface {
	Show()
}

/*
Module Layer
*/

type Apple struct {
}

func (apple *Apple) Show() {
	fmt.Println("Apple")
}

type Banana struct {
}

func (banana *Banana) Show() {
	fmt.Println("Banana")
}

type Orange struct {
}

func (orange *Orange) Show() {
	fmt.Println("Orange")
}

/*
Factory
*/

type Factory struct {
}

func (factory *Factory) CreateFruit(name string) Fruit {
	// parent pointer
	var fruit Fruit
	// points to children obj = polymorphism
	switch name {
	case "Apple":
		fruit = &Apple{}
	case "Banana":
		fruit = &Banana{}
	case "Orange":
		fruit = &Orange{}
	}
	return fruit
}

/*
Business Logic Layer
*/
func main() {
	factory := &Factory{}
	// apple/banana/orange is Fruit type, interface-oriented
	apple := factory.CreateFruit("Apple")
	banana := factory.CreateFruit("Banana")
	orange := factory.CreateFruit("Orange")
	apple.Show()
	banana.Show()
	orange.Show()
}
