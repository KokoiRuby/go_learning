package main

import "fmt"

type AbstractBanker interface {
	DoBusiness()
}

type DepositBanker struct {
}

func (d *DepositBanker) DoBusiness() {
	fmt.Println("deposit")
}

type TransferBanker struct {
}

func (d *TransferBanker) DoBusiness() {
	fmt.Println("transfer")
}

type PayBanker struct {
}

func (d *PayBanker) DoBusiness() {
	fmt.Println("pay")
}

/*
ShareBanker (+)
*/
type ShareBanker struct {
}

func (d *ShareBanker) DoBusiness() {
	fmt.Println("share")
}

/*
BankBusiness (+) encapsulates interface as parent pointer
*/
func BankBusiness(ab AbstractBanker) {
	// polymorphism
	ab.DoBusiness()
}

func main() {
	db := DepositBanker{}
	db.DoBusiness()

	tb := TransferBanker{}
	tb.DoBusiness()

	pb := PayBanker{}
	pb.DoBusiness()

	// (+)
	sb := ShareBanker{}
	sb.DoBusiness()

	// polymorphism, pass children obj
	BankBusiness(&db)
	BankBusiness(&tb)
	BankBusiness(&pb)
	BankBusiness(&sb)
}
