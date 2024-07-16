package main

import "fmt"

type Point1 struct {
	x, y int
}

type Rectangle1 struct {
	a, b Point1
}

func main() {
	r1 := Rectangle1{Point1{1, 2}, Point1{3, 4}}
	fmt.Println(&r1.a.x, &r1.a.y)
	fmt.Println(&r1.b.x, &r1.b.y)

}
