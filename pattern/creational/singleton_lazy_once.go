package main

import (
	"fmt"
	"sync"
)

// private class
type singleton struct{}

// points to nil at the beginning
var instance *singleton
var once sync.Once

// public to get instance (lazy)
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
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
