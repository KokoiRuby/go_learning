package main

import "fmt"

/*
Abstraction
*/

type Fruit interface {
	Show()
}

type AbstractFactory interface {
	CreateFruit() Fruit
}

/*
Module
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

type AppleFactory struct {
	AbstractFactory
}

func (factory *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = &Apple{} // parent pointer to child obj
	return fruit
}

type BananaFactory struct {
	AbstractFactory
}

func (factory *BananaFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = &Banana{} // parent pointer to child obj
	return fruit
}

type OrangeFactory struct {
	AbstractFactory
}

func (factory *OrangeFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = &Orange{} // parent pointer to child obj
	return fruit
}

/*
Business Logic
*/
func main() {
	var appleFactory AbstractFactory // abstraction-oriented
	appleFactory = &AppleFactory{}
	var apple Fruit // abstraction-oriented
	apple = appleFactory.CreateFruit()
	apple.Show()

	var bananaFactory AbstractFactory // abstraction-oriented
	bananaFactory = &BananaFactory{}
	var banana Fruit // abstraction-oriented
	banana = bananaFactory.CreateFruit()
	banana.Show()

	var orangeFactory AbstractFactory // abstraction-oriented
	orangeFactory = &OrangeFactory{}
	var orange Fruit // abstraction-oriented
	orange = orangeFactory.CreateFruit()
	orange.Show()
}
