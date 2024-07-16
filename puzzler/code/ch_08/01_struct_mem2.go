package main

import "fmt"

type Point2 struct {
	x, y int
}

type Rectangle2 struct {
	a, b *Point2
}

func main() {
	r2 := Rectangle2{&Point2{5, 6}, &Point2{7, 8}}
	fmt.Println(&r2.a, &r2.b)
	fmt.Println(r2.a, r2.b)
	fmt.Println(*r2.a, *r2.b)

}
