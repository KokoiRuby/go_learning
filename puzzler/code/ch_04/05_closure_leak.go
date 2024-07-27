package main

import (
	"fmt"
	"time"
)

func main() {
	var data []int

	closure := func() {
		fmt.Println(data)
	}

	for i := 0; i < 1000; i++ {
		go func() {
			for {
				time.Sleep(1 * time.Second)
			}
		}()
	}

	closure()

	// data no longer needed, but still kept by closure, not able to be GC
}
