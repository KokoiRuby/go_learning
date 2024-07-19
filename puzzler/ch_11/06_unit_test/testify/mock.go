package main

import "fmt"

type Printer interface {
	Print(paperSize string, content string) error
}

type MyPrinter struct {
}

func (p MyPrinter) Print(paperSize string, content string) error {
	fmt.Println("Printing:", content, ", with paper size:", paperSize)
	return nil
}

type Computer struct {
	Pr Printer
}

func (c *Computer) PrintA4(content string) error {
	c.Pr.Print("A4", content)
	return nil
}
