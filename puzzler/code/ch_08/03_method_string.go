package main

import "fmt"

type Bird struct {
	name string
	age  int
}

func (b *Bird) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", b.name, b.age)
}

func main() {
	bird := &Bird{
		name: "Alice",
		age:  25,
	}

	fmt.Println(bird)
}
