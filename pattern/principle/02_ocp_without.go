package main

import "fmt"

type Banker struct {
}

func (b *Banker) Deposit() {
	fmt.Println("Banker deposit")
}

func (b *Banker) Transfer() {
	fmt.Println("Banker transfer")
}

func (b *Banker) Pay() {
	fmt.Println("Banker pay")
}

/*
Share (+) against OCP
*/
func (b *Banker) Share() {
	fmt.Println("Banker share")
}

func main() {
	banker := Banker{}
	banker.Deposit()
	banker.Transfer()
	banker.Pay()
	// (+)
	banker.Share()
}
