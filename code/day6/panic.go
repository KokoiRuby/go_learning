package main

import "fmt"

func panicFunc() {
	panic("!")
}

func main() {
	panicFunc()
	fmt.Println("Last word is:...")
}
