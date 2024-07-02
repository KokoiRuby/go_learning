package main

import (
	"fmt"
	"strconv"
	"time"
)

func strToNum() {
	for i := 1; i <= 10; i++ {
		fmt.Println("[goroutine]", strconv.Itoa(i))
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go strToNum()
	fmt.Println(" ")
	for i := 1; i <= 10; i++ {
		fmt.Println("[main]", strconv.Itoa(i))
		time.Sleep(time.Second * 1)
	}
}
