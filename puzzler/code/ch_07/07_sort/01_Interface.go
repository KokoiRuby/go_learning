package main

import (
	"fmt"
	"sort"
)

func main() {
	// int
	var intSl sort.IntSlice
	intSl = append(intSl, []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}...)

	fmt.Println(intSl)
	fmt.Println(intSl.Len())
	fmt.Println(intSl.Less(0, 1))
	intSl.Swap(0, 1)
	fmt.Println(intSl)
	fmt.Println(intSl.Search(5))

	sort.Ints(intSl)
	fmt.Println(intSl)
	fmt.Println(sort.IsSorted(intSl))
	fmt.Println(sort.SearchInts(intSl, 5))
	sort.Sort(sort.Reverse(intSl))
	fmt.Println(intSl)
	fmt.Println()

	// float
	var floatSl sort.Float64Slice
	floatSl = []float64{3.14, 1.618, 2.718, 1.414, 2.236}

	fmt.Println(floatSl)
	fmt.Println(floatSl.Len())
	// math.isNaN(collection[i])
	fmt.Println(floatSl.Less(0, 1))
	floatSl.Swap(0, 1)
	fmt.Println(floatSl)
	fmt.Println(floatSl.Search(2.718))

	sort.Float64s(floatSl)
	fmt.Println(floatSl)
	fmt.Println(sort.Float64sAreSorted(floatSl))
	fmt.Println(sort.SearchFloat64s(floatSl, 2.236))
	sort.Sort(sort.Reverse(floatSl))
	fmt.Println(floatSl)
	fmt.Println()

	// string
	var stringSl sort.StringSlice
	stringSl = []string{"apple", "banana", "cherry", "date", "fig"}

	fmt.Println(stringSl)
	fmt.Println(stringSl.Len())
	fmt.Println(stringSl.Less(0, 1))
	stringSl.Swap(0, 1)
	fmt.Println(stringSl)
	fmt.Println(stringSl.Search("apple"))

	sort.Strings(stringSl)
	fmt.Println(stringSl)
	fmt.Println(sort.StringsAreSorted(stringSl))
	fmt.Println(sort.SearchStrings(stringSl, "date"))
	sort.Sort(sort.Reverse(stringSl))
	fmt.Println(stringSl)

}
