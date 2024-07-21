package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// private class
type singleton struct{}

// points to nil at the beginning
var instance *singleton
var lock sync.Mutex
var initialized uint32

// public to get instance (lazy)
func GetInstance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = &singleton{}
		atomic.AddUint32(&initialized, 1)
		return instance
	}
	return instance
}

func (s *singleton) DoSomething() {
	fmt.Println("do something")
}

func main() {
	s1 := GetInstance()
	s1.DoSomething()

	s2 := GetInstance()
	if s1 == s2 {
		fmt.Println("equal")
	}

}
