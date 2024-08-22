package main

import (
	"fmt"
	"sort"
)

// matrix

type matrix [][]int

func (m matrix) Len() int {
	return len(m)
}

func (m matrix) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m matrix) Less(i, j int) bool {
	return m[i][1] < m[j][1]
}

// map

type mapSl []map[string]int

func (msl mapSl) Len() int {
	return len(msl)
}
func (msl mapSl) Swap(i, j int) {
	msl[i], msl[j] = msl[j], msl[i]
}
func (msl mapSl) Less(i, j int) bool {
	return msl[i]["a"] < msl[j]["a"]
}

func main() {
	m := matrix{
		{1, 4},
		{9, 3},
		{7, 5},
	}

	fmt.Println(m)
	fmt.Println(len(m))
	m.Swap(0, 1)
	fmt.Println(m)
	sort.Sort(m)
	fmt.Println()
	fmt.Println()

	msl := mapSl{
		{"a": 4, "b": 12},
		{"a": 3, "b": 11},
		{"a": 5, "b": 10},
	}

	fmt.Println(msl)
	sort.Sort(msl)
	fmt.Println(msl)
	fmt.Println(msl.Len())
	msl.Swap(0, 1)
	fmt.Println(msl)

}
