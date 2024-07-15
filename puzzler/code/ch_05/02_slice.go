package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[1:3]
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(sl)
	fmt.Println(cap(sl))
	fmt.Println(len(sl))

	// modify
	arr[2] = 10
	fmt.Println(sl)

	// make
	var sl2 []int = make([]int, 5, 10)
	fmt.Println(sl2)

	// iter1
	for i := 0; i < len(sl2); i++ {
		println(sl2[i])
	}

	// iter2
	for i, v := range sl2 {
		println(i, v)
	}

	// nil
	var nilSlice []int
	if nilSlice == nil {
		fmt.Println("nilSlice is nil.")
	}

}
