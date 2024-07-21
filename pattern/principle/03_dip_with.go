package main

import "fmt"

// abs layer

type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// impl layer

type Benz struct {
}

func (b *Benz) Run() {
	fmt.Println("Benz Run")
}

type BMW struct {
}

func (b *BMW) Run() {
	fmt.Println("BMW Run")
}

// (+)

type BYD struct {
}

func (b *BYD) Run() {
	fmt.Println("BYD Run")
}

type Alice struct {
}

func (a *Alice) Drive(car Car) {
	fmt.Println("Alice Drive")
	car.Run()
}

type Bob struct {
}

func (b *Bob) Drive(car Car) {
	fmt.Println("Bob Drive")
	car.Run()
}

// (+)

type Cindy struct {
}

func (c *Cindy) Drive(car Car) {
	fmt.Println("Cindy Drive")
	car.Run()
}

// business layer

func main() {
	var benz Car // dip, dep on abs
	benz = new(Benz)
	var a Driver // dip, dep on abs
	a = new(Alice)
	a.Drive(benz)

	var bmw Car
	bmw = new(BMW)
	var b Driver
	b = new(Bob)
	b.Drive(bmw)

	// (+)
	a.Drive(bmw)
	b.Drive(benz)

	var byd Car
	byd = new(BYD)
	var c Driver
	c = new(Cindy)
	c.Drive(byd)

}
