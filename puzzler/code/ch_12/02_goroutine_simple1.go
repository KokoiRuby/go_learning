package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(0))
	go doSomething()
	// ++ otherwise, main exit before goroutine finished.
	time.Sleep(time.Second)
	fmt.Println("done")
}

func doSomething() {
	fmt.Println("doSomething called")
}
