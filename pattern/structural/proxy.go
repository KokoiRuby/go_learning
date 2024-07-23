package main

import "fmt"

type Good struct {
	Kind string
	Fact bool
}

/*
Abstract
*/

type Shopping interface {
	Buy(good *Good)
}

/*
Impl
*/

type KoreaShopping struct{}

func (ks KoreaShopping) Buy(good *Good) {
	fmt.Println("korea shopping", good.Kind)
}

type USShopping struct{}

func (us USShopping) Buy(good *Good) {
	fmt.Println("us shopping", good.Kind)
}

type UKShopping struct{}

func (uk UKShopping) Buy(good *Good) {
	fmt.Println("uk shopping", good.Kind)
}

type OverseasProxy struct {
	shopping Shopping // abs
}

func NewOverseasProxy(shopping Shopping) *OverseasProxy {
	return &OverseasProxy{shopping}
}

func (op OverseasProxy) Buy(good *Good) {
	// ++
	if op.isFact(good) {
		op.Check()
		op.shopping.Buy(good)
	} else {
		fmt.Println("Fake")
	}
}

func (op OverseasProxy) isFact(good *Good) bool {
	return good.Fact
}

func (op OverseasProxy) Check() {
	fmt.Println("Checked")
}

/*
Busi
*/
func main() {
	g1 := &Good{
		Kind: "bag",
		Fact: true,
	}

	g2 := &Good{
		Kind: "tag",
		Fact: false,
	}

	var shopping Shopping
	shopping = USShopping{}

	var op Shopping
	op = NewOverseasProxy(shopping)
	op.Buy(g1)

	shopping = UKShopping{}
	op = NewOverseasProxy(shopping)
	op.Buy(g2)

}
