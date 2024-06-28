package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		fmt.Println(rand.Intn(100))
		time.Sleep(time.Second * 1)
	}

}
