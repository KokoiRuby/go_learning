package main

import "fmt"

func main() {
re:
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		for j := 11; j <= 13; j++ {
			fmt.Println(j)
			continue re
		}
	}
}
