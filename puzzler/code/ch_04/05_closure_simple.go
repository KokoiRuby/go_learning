package main

import "fmt"

func getSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nxt := getSeq()

	fmt.Println(nxt())
	fmt.Println(nxt())
	fmt.Println(nxt())

	nxtNew := getSeq()
	fmt.Println(nxtNew())
	fmt.Println(nxtNew())
	fmt.Println(nxtNew())
}
