package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	rwMutex   sync.RWMutex
	sharedVar int
)

func read() {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	fmt.Println("Read value: ", sharedVar)
	time.Sleep(time.Microsecond * 1000)
}

func write(val int) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	sharedVar = val
	fmt.Println("Write value: ", sharedVar)
	time.Sleep(time.Microsecond * 1000)
}

func main() {
	for i := 0; i < 5; i++ {
		go read()
		go write(i)
	}

	time.Sleep(time.Microsecond * 8000)
}
