package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d - addr is: %v\n", arr[i], &arr[i])
	}

	for _, v := range arr {
		fmt.Printf("%d - addr is: %v\n", v, &v)
	}
}
