package main

import "fmt"

type Cat struct {
	Name  string
	Age   int
	color string
	ptr   *int
	sl    []int
	m     map[string]string
}

func main() {
	var cat Cat
	fmt.Println(cat)
	fmt.Println(&cat)
	cat.Name = "Kitty"
	cat.Age = 18
	cat.color = "White"
	fmt.Println(cat)
	fmt.Println(cat.Name)
	fmt.Println(cat.Age)
	fmt.Println(cat.color)
	fmt.Println(&cat)

	num := 1
	cat.ptr = &num
	fmt.Println(cat.ptr)
	cat.sl = make([]int, 3)
	cat.sl = []int{1, 2, 3}
	fmt.Println(cat.sl)
	cat.m = make(map[string]string)
	cat.m["key"] = "value"
	fmt.Println(cat.m)
}
