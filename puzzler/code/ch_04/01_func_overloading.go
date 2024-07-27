package main

import (
	"fmt"
	"math"
)

// 1. interface

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

// 2. by variadic

func Sum(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}

// 3. by func option

type Config struct {
	Option1 string
	Option2 int
}

type Option func(*Config) // type redef

func WithOption1(value string) Option {
	return func(config *Config) {
		config.Option1 = value
	}
}

func WithOption2(value int) Option {
	return func(config *Config) {
		config.Option2 = value
	}
}

func NewConfig(opts ...Option) *Config {
	config := &Config{}
	for _, opt := range opts {
		opt(config)
	}
	return config
}

func main() {
	circle := &Circle{Radius: 1}
	rectangle := &Rectangle{Width: 10, Height: 20}
	fmt.Println(circle.Area())
	fmt.Println(rectangle.Area())

	fmt.Println(Sum(1))
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(1, 2, 3, 4, 5))

	//config := NewConfig(WithOption1("value1"), WithOption2(42))
}
