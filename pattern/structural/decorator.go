package main

import "fmt"

/*
Abs
*/

type Phone interface {
	Show()
}

type Decorator struct {
	phone Phone
}

func (d Decorator) Show() {}

/*
Impl - component
*/

type Apple struct{}

func (a Apple) Show() {
	fmt.Println("Apple")
}

type Xiaomi struct{}

func (xiaomi Xiaomi) Show() {
	fmt.Println("Xiaomi")
}

/*
Impl - decorator
*/

type screenProtectorDecorator struct {
	Decorator
}

func (spd *screenProtectorDecorator) Show() {
	spd.phone.Show()
	// +
	fmt.Println("++ screen protector")
}

func NewScreenProtector(phone Phone) Phone {
	return &screenProtectorDecorator{Decorator{phone}}
}

type shellDecorator struct {
	Decorator
}

func (sd shellDecorator) Show() {
	sd.phone.Show()
	// +
	fmt.Println("++ shell")
}

func NewShellDecorator(phone Phone) Phone {
	return &shellDecorator{Decorator{phone}}
}

/*
Business Logic
*/
func main() {
	var iphone Phone
	iphone = &Apple{}
	iphone.Show()

	var iphoneWithShell Phone
	iphoneWithShell = NewShellDecorator(iphone)
	iphoneWithShell.Show()

	var iphoneWithShellWithScreenProtector Phone
	iphoneWithShellWithScreenProtector = NewScreenProtector(iphoneWithShell)
	iphoneWithShellWithScreenProtector.Show()
}
