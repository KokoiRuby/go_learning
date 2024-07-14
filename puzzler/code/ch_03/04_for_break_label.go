package main

import (
	"fmt"
)

func main() {
outer:
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if i*j >= 12 {
				fmt.Println(i, j)
				break outer
			}
		}
	}
}
