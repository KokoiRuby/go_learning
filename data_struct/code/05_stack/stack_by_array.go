package _5_stack

import (
	"errors"
	"fmt"
)

type ArrayStack struct {
	data []interface{}
	top  int
}

func NewArrayStack(capacity uint) *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, capacity, capacity),
		top:  0,
	}
}

func (s *ArrayStack) IsEmpty() bool {
	if s.top == 0 {
		return true
	}
	return false
}

func (s *ArrayStack) Push(v interface{}) error {
	if s.top == cap(s.data) {
		return errors.New("stack is full")
	}
	s.data[s.top] = v
	s.top++
	return nil
}

func (s *ArrayStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	s.top--
	v := s.data[s.top]
	return v, nil
}

func (s *ArrayStack) Flush() {
	for i, _ := range s.data {
		s.data[i] = nil
	}
	s.top = -1
}

func (s *ArrayStack) Print() {
	for i, _ := range s.data {
		fmt.Printf("%v ", s.data[i])
	}
}
