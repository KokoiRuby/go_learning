package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	now := time.Now()
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	end := time.Now().UnixNano()
	fmt.Println("cost:", end-start, "(ns)")
}
