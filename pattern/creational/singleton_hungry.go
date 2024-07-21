package main

import "fmt"

// private class
type singleton struct{}

// points to the only instance (hungry)
var instance *singleton = &singleton{}

// public to get instance
func GetInstance() *singleton {
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
