package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Usb2 interface {
	Start()
	Stop()
	Test()
}

type Phone struct {
}

type Camera struct {
}

func (p Phone) Start() {
	fmt.Println("Starting...")
}

func (p Phone) Stop() {
	fmt.Println("Stopping...")
}

func (c Camera) Start() {
	fmt.Println("Starting...")
}

func (c Camera) Stop() {
	fmt.Println("Stopping...")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)
}
