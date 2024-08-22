package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type Persons []Person

func (p Persons) Len() int {
	return len(p)
}

func (p Persons) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	p := &Persons{
		Person{
			"Alice",
			1,
		},
		Person{
			"Bob",
			4,
		},
		Person{
			"Cindy",
			2,
		},
		Person{
			"Dan",
			3,
		},
	}
	fmt.Println(p.Len())
	sort.Sort(p)
	fmt.Println(p)
	p.Swap(0, 1)
	fmt.Println(p)
}
