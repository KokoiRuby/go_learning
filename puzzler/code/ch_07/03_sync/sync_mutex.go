package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var wg sync.WaitGroup
var count int

func f1() {
	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()
	count++
	fmt.Println("[f1] count is: ", count)
}

func f2() {
	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()
	count++
	fmt.Println("[f2] count is: ", count)
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go f1()
		go f2()
	}
	wg.Wait()
}
