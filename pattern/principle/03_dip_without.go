package main

import "fmt"

// Cars

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

// Drivers

type Alice struct {
}

func (a *Alice) DriveBenz(benz *Benz) {
	fmt.Println("Alice Drive Benz")
	benz.Run()
}

// DriveBMW (+)
func (a *Alice) DriveBMW(bmw *BMW) {
	fmt.Println("Alice Drive BMW")
	bmw.Run()
}

type Bob struct {
}

func (b *Bob) DriveBMW(bmw *BMW) {
	fmt.Println("Bob Drive BMW")
	bmw.Run()
}

// DriveBenz (+)
func (b *Bob) DriveBenz(benz *Benz) {
	fmt.Println("Alice Drive Benz")
	benz.Run()
}

func main() {
	benz := &Benz{}
	alice := &Alice{}
	alice.DriveBenz(benz)

	bmw := &BMW{}
	bob := &Bob{}
	bob.DriveBMW(bmw)

	// (+)
	alice.DriveBMW(bmw)
	bob.DriveBenz(benz)
}
