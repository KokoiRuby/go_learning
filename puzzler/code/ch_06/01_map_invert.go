package main

import "fmt"

func main() {
	originalMap := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	invertedMap := make(map[int]string)

	for key, value := range originalMap {
		invertedMap[value] = key
	}

	for key, value := range invertedMap {
		fmt.Printf("%d: %s\n", key, value)
	}
}
