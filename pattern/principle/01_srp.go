package main

import "fmt"

/*
Clothes has 2 responsibilities which is against SRP
*/
type Clothes struct {
}

func (c *Clothes) OnWork() {
	fmt.Println("Working Clothes")
}

func (c *Clothes) OnShop() {
	fmt.Println("Shopping Clothes")
}

/*
ClothesWork for work only, SRP
*/
type ClothesWork struct {
}

func (c *ClothesWork) Style() {
	fmt.Println("Working Clothes")
}

/*
ClothesShop for shop only, SRP
*/
type ClothesShop struct {
}

func (c *ClothesShop) Style() {
	fmt.Println("Shopping Clothes")
}

func main() {
	c := Clothes{}
	c.OnWork()
	c.OnShop()

	cw := ClothesWork{}
	cw.Style()
	cs := ClothesShop{}
	cs.Style()
}
