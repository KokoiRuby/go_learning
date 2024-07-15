package main

import "fmt"

func traceFunc(name string) func() {
	fmt.Println("Entering", name)
	return func() {
		fmt.Println("Exiting", name)
	}
}

func myFunction() {
	// defer func() {return fmt.Println("Exiting", myFunction)}
	defer traceFunc("myFunction")()
	fmt.Println("Inside myFunction")
}

func main() {
	myFunction()
}
