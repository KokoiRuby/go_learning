package main

import "fmt"

/*
Abstraction - Factory
*/

type AbsFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
	CreateOrange() AbstractOrange
}

/*
Abstraction - Product
*/

type AbstractApple interface {
	ShowApple()
}

type AbstractBanana interface {
	ShowBanana()
}

type AbstractOrange interface {
	ShowOrange()
}

/*
Implementation - China product group
*/

type ChineseApple struct{}

func (ca *ChineseApple) ShowApple() {
	fmt.Println("Chinese Apple")
}

type ChineseBanana struct{}

func (cb *ChineseBanana) ShowBanana() {
	fmt.Println("Chinese Banana")
}

type ChineseOrange struct{}

func (co *ChineseOrange) ShowOrange() {
	fmt.Println("Chinese Orange")
}

type ChineseFactory struct{}

func (cf ChineseFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = &ChineseApple{}
	return apple
}

func (cf ChineseFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = &ChineseBanana{}
	return banana
}

func (cf ChineseFactory) CreateOrange() AbstractOrange {
	var orange AbstractOrange
	orange = &ChineseOrange{}
	return orange
}

/*
Implementation - Japan product group
*/

type JapaneseApple struct{}

func (ja *JapaneseApple) ShowApple() {
	fmt.Println("Japanese Apple")
}

type JapaneseBanana struct{}

func (jb *JapaneseBanana) ShowBanana() {
	fmt.Println("Japanese Banana")
}

type JapaneseOrange struct{}

func (jo *JapaneseOrange) ShowOrange() {
	fmt.Println("Japanese Orange")
}

type JapaneseFactory struct{}

func (jf JapaneseFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = &JapaneseApple{}
	return apple
}

func (jf JapaneseFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = &JapaneseBanana{}
	return banana
}

func (jf JapaneseFactory) CreateOrange() AbstractOrange {
	var orange AbstractOrange
	orange = &JapaneseOrange{}
	return orange
}

/*
Business Logic
*/
func main() {
	var cf AbsFactory
	cf = &ChineseFactory{}

	var ca AbstractApple
	var cb AbstractBanana
	var co AbstractOrange
	ca = cf.CreateApple()
	ca.ShowApple()
	cb = cf.CreateBanana()
	cb.ShowBanana()
	co = cf.CreateOrange()
	co.ShowOrange()

	var jf AbsFactory
	jf = &JapaneseFactory{}

	var ja AbstractApple
	var jb AbstractBanana
	var jo AbstractOrange
	ja = jf.CreateApple()
	ja.ShowApple()
	jb = jf.CreateBanana()
	jb.ShowBanana()
	jo = jf.CreateOrange()
	jo.ShowOrange()

}
