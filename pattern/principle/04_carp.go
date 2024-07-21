package main

import "fmt"

type Cat struct {
}

func (c *Cat) Eat() {
	fmt.Println("Cat is eating")
}

// (+) by inheritance

type CatB struct {
	Cat
}

func (c *CatB) Sleep() {
	fmt.Println("Cat is sleeping")
}

// (+) by combination partial

type CatC struct {
	// c *Cat
}

func (cc *CatC) Eat(c *Cat) {
	// only eat need to combine with Cat
	c.Eat()
}

func (cc *CatC) Sleep() {
	fmt.Println("Cat is sleeping")
}

func main() {
	c := &Cat{}
	c.Eat()

	// inheritance
	cb := &CatB{}
	cb.Eat()
	cb.Sleep()

	/// combination
	cc := &CatC{}
	cc.Eat(c)
	cc.Sleep()
}
