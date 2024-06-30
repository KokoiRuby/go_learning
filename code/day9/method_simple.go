package main

import "fmt"

type A struct {
	Age int
}

func (a A) test() {
	fmt.Println(a.Age)
}

func main() {
	var a A
	a.Age = 20
	a.test()
}
