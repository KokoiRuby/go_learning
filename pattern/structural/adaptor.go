package main

import "fmt"

/*
Client Interface
*/

type V5 interface {
	Use5V()
}

/*
Client
*/

type Phone struct {
	v V5
}

func NewPhone(v V5) *Phone {
	return &Phone{v: v}
}

func (p *Phone) Charge() {
	fmt.Println("Charging")
	p.v.Use5V()
}

/*
Adaptee
*/

type V220 struct {
}

func (v *V220) Use220V() {
	fmt.Println("Using V220 to charge")
}

/*
Adaptor
*/

type Adapter struct {
	v *V220
}

func NewAdapter(v *V220) *Adapter {
	return &Adapter{v: v}
}

func (a *Adapter) Use5V() {
	fmt.Println("Using Adapter to charge")
	a.v.Use220V()
}

/*
Busi
*/

func main() {
	iphone := NewPhone(NewAdapter(new(V220)))
	iphone.Charge()
}
