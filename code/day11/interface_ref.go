package main

import "fmt"

type Usb interface {
	Say()
}

type Stu struct {
}

func (this *Stu) Say() {
	fmt.Println("Say something...")
}

func main() {
	stu := Stu{}
	// var u Usb = stu
	var u Usb = &stu
	u.Say()
}
