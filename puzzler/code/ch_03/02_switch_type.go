package main

import "fmt"

func main() {
	var x interface{}

	switch x.(type) {
	case nil:
		fmt.Println("type of x is nil.")
	case int:
		fmt.Println("type of x is int.")
	case float32:
		fmt.Println("type of x is float32.")
	case func(int) float64:
		fmt.Println("type of x is int.")
	case bool, string:
		fmt.Println("type of x is bool or string.")
	default:
		fmt.Println("Unknown.")
	}

}
