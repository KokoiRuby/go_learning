package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	name string
}

func (p Phone) Start() {
	fmt.Println("Starting...")
}

func (p Phone) Stop() {
	fmt.Println("Stopping...")
}

type Camera struct {
	name string
}

func (c Camera) Start() {
	fmt.Println("Starting...")
}

func (c Camera) Stop() {
	fmt.Println("Stopping...")
}

func main() {
	var usbArr [2]Usb
	usbArr[0] = Phone{"iPhone"}
	usbArr[1] = Camera{"Nikon"}

	for _, elem := range usbArr {
		elem.Start()
		elem.Stop()
	}
}
