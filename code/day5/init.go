package main

import "fmt"

const pi = 3.14

var globalVar = genGlobalVar()

func genGlobalVar() int {
	fmt.Println("gen global var")
	return 1
}

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
}
