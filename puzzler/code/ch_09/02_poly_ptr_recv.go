package main

import "fmt"

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c *cat) move() {
	fmt.Println("Moving...")
}

func (c *cat) eat(food string) {
	fmt.Println("Eating", food)
}

func main() {
	var cat = &cat{"Tom", 4}
	var animal = cat
	animal.move()
	animal.eat("Fish")
}
