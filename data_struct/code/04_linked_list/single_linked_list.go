package _4_linked_list

import (
	"errors"
	"fmt"
)

type Node struct {
	next *Node
	val  interface{}
}

type SingleLinkedList struct {
	head   *Node
	length uint
}

func NewNode(v interface{}) *Node {
	return &Node{val: v}
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{
		head:   NewNode(nil),
		length: 0,
	}
}

func (sll *SingleLinkedList) GetHead() *Node {
	return sll.head
}

func (sll *SingleLinkedList) GetTail() *Node {
	cur := sll.head
	for cur.next != nil {
		cur = cur.next
	}
	return cur
}

func (sll *SingleLinkedList) Append(v interface{}) {
	if sll.head.next == nil {
		sll.head.next = NewNode(v)
	} else {
		tail := sll.GetTail()
		tail.next = NewNode(v)
	}
	sll.length++
}

func (sll *SingleLinkedList) FindByIndex(idx uint) (prev *Node, curr *Node, err error) {
	if idx >= sll.length {
		return nil, nil, errors.New("out of range")
	}
	// iter from head
	pre := sll.head
	cur := sll.head.next
	for i := uint(0); i < idx; i++ {
		pre = cur
		cur = cur.next
	}
	return pre, cur, nil
}

func (sll *SingleLinkedList) InsertAfter(idx uint, v interface{}) (bool, error) {
	if idx >= sll.length {
		return false, errors.New("out of range")
	}
	_, cur, _ := sll.FindByIndex(idx)
	newNode := NewNode(v)
	newNode.next = cur.next
	cur.next = newNode
	sll.length++
	return true, nil
}

func (sll *SingleLinkedList) InsertBefore(idx uint, v interface{}) (bool, error) {
	if idx >= sll.length {
		return false, errors.New("out of range")
	}
	pre, cur, _ := sll.FindByIndex(idx)
	newNode := NewNode(v)
	newNode.next = cur
	pre.next = newNode
	sll.length++
	return true, nil
}

func (sll *SingleLinkedList) Remove(idx uint) interface{} {
	if idx >= sll.length {
		return nil
	}
	pre, cur, _ := sll.FindByIndex(idx)
	pre.next = cur.next
	v := cur.val
	cur = nil
	sll.length--
	return v
}

func (sll *SingleLinkedList) Print() {
	cur := sll.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.val)
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}
