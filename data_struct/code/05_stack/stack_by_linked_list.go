package _5_stack

import "errors"

type Node struct {
	next *Node
	val  interface{}
}

type LinkedListStack struct {
	top      *Node
	length   int
	capacity int
}

func NewLinkedListStack(capacity int) *LinkedListStack {
	return &LinkedListStack{
		top:      nil,
		length:   0,
		capacity: capacity,
	}
}

func (s *LinkedListStack) Push(val interface{}) error {
	if s.length == s.capacity {
		return errors.New("stack is full")
	}
	s.top = &Node{next: s.top, val: val}
	s.length++
	return nil
}

func (s *LinkedListStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	v := s.top.val
	s.top = s.top.next
	s.length--
	return v, nil
}

func (s *LinkedListStack) IsEmpty() bool {
	return s.length == 0
}

func (s *LinkedListStack) Flush() {
	s.top = nil
}
