package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Circle struct {
	radius float32
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func main() {
	var shaper Shaper
	square := &Square{5}
	shaper = square

	if t, ok := shaper.(*Square); ok {
		fmt.Printf("Type is %T\n", t)
	}
	if t, ok := shaper.(*Circle); ok {
		fmt.Printf("Type is %T\n", t)
	}

	// type swtich
	switch t := shaper.(type) {
	case *Square:
		fmt.Printf("Type is %T", t)
	case *Circle:
		fmt.Printf("Type is %T", t)
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}

}
