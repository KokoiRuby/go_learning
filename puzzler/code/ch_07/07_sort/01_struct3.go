package main

import (
	"fmt"
	"sort"
)

type Fruit struct {
	Name   string
	Weight int
}

type Fruits []Fruit

func (f Fruits) Len() int {
	return len(f)
}

func (f Fruits) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// enCap into ByName

type ByName struct{ Fruits }

func (b ByName) Less(i, j int) bool {
	return b.Fruits[i].Name < b.Fruits[j].Name
}

// enCap to Weight

type ByWeight struct{ Fruits }

func (b ByWeight) Less(i, j int) bool {
	return b.Fruits[i].Weight < b.Fruits[j].Weight
}

func main() {
	fruits := Fruits{
		Fruit{"apple", 10},
		Fruit{"peach", 30},
		Fruit{"pear", 20},
		Fruit{"banana", 50},
		Fruit{"pineapple", 40},
		Fruit{"grape", 70},
		Fruit{"durian", 60},
	}

	sort.Sort(ByName{fruits})
	fmt.Println(fruits)

	sort.Sort(ByWeight{fruits})
	fmt.Println(fruits)
}
