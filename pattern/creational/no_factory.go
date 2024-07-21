package main

import "fmt"

type Fruit struct {
}

func (f *Fruit) show(name string) {
	if name == "apple" {
		fmt.Println("I'm apple")
	} else if name == "orange" {
		fmt.Println("I'm orange")
	} else if name == "banana" {
		fmt.Println("I'm banana")
	}
}

func NewFruit(name string) *Fruit {
	fruit := new(Fruit)
	if name == "apple" {
		fmt.Println("new apple")
	} else if name == "orange" {
		fmt.Println("new orange")
	} else if name == "banana" {
		fmt.Println("new banana")
	}
	return fruit
}

func main() {
	apple := NewFruit("apple")
	apple.show("apple")

	banana := NewFruit("banana")
	banana.show("banana")

	orange := NewFruit("orange")
	orange.show("orange")

}
