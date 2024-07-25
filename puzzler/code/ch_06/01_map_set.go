package main

import (
	"fmt"
	"sync"
)

type MapSet struct {
	m       map[interface{}]struct{}
	rwMutex sync.RWMutex
}

func NewMapSet() *MapSet {
	return &MapSet{m: make(map[interface{}]struct{}), rwMutex: sync.RWMutex{}}
}

func (s *MapSet) Add(key interface{}) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	s.m[key] = struct{}{}
}

func (s *MapSet) Remove(key interface{}) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	delete(s.m, key)
}

func (s *MapSet) Contains(key interface{}) bool {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()

	_, ok := s.m[key]
	return ok
}

func main() {
	ms := NewMapSet()
	ms.Add("A")
	ms.Add("B")
	ms.Add("C")
	ms.Add("C")

	fmt.Println(ms.Contains("A"))
	fmt.Println(ms.Contains("D"))

	ms.Remove("A")
	fmt.Println(ms.Contains("A"))
}
