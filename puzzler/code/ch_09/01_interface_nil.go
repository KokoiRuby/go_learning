package main

import "fmt"

func main() {
	var a interface{}
	a = 20
	fmt.Println(a)
	a = "Hello"
	fmt.Println(a)
	a = true
	fmt.Println(a)

	var m = make(map[string]interface{})
	m["key1"] = 20
	m["key2"] = "Hello"
	m["key3"] = true
	fmt.Println(m)

	var sl = []interface{}{1, 2, 3, 4, 5, "hello", true}
	fmt.Println(sl)

}
