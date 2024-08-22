package main

import (
	"fmt"
	"sort"
)

type Cat struct {
	Name string
	Age  int
}

type CatWrapper struct {
	Cats []Cat
	by   func(p1, p2 *Cat) bool
}

func (cw *CatWrapper) Len() int {
	return len(cw.Cats)
}

// Less is dynamic
func (cw *CatWrapper) Less(i, j int) bool {
	return cw.by(&cw.Cats[i], &cw.Cats[j])
}

func (cw *CatWrapper) Swap(i, j int) {
	cw.Cats[i], cw.Cats[j] = cw.Cats[j], cw.Cats[i]
}

type SortBy func(p, q *Cat) bool

func SortCats(cats []Cat, by SortBy) {
	sort.Sort(&CatWrapper{cats, by})
}

func main() {
	cats := []Cat{
		Cat{
			"Alice",
			1,
		},
		Cat{
			"Bob",
			4,
		},
		Cat{
			"Cindy",
			2,
		},
		Cat{
			"Dan",
			3,
		},
	}

	fmt.Println(cats)

	catsWrapper := &CatWrapper{
		Cats: cats,
		by: func(p1, p2 *Cat) bool {
			return p1.Age < p2.Age
			//return p1.Name < p2.Name
		},
	}

	fmt.Println(catsWrapper.Len())
	catsWrapper.Swap(0, 1)
	fmt.Println(cats)

	SortCats(cats, func(p, q *Cat) bool {
		return p.Age < q.Age
	})
	fmt.Println(cats)

}
