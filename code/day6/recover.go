package main

import "fmt"

func panicFunc() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("!")
}

func main() {
	panicFunc()
	fmt.Println("Last word is:...")
}
