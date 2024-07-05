package _4_linked_list

import (
	"errors"
	"fmt"
)

type NodeCD struct {
	prev *NodeCD
	next *NodeCD
	val  interface{}
}

type CircularDoublyLinkedList struct {
	head   *NodeCD
	length uint
}

func NewNodeCD(v interface{}) *NodeCD {
	return &NodeCD{
		val: v,
	}
}

func NewCircularDoublyLinkedList() *CircularDoublyLinkedList {
	return &CircularDoublyLinkedList{
		head:   NewNodeCD(nil),
		length: 0,
	}
}

func (cdll *CircularDoublyLinkedList) GetHead() *NodeCD {
	return cdll.head
}

func (cdll *CircularDoublyLinkedList) GetTail() *NodeCD {
	cur := cdll.head
	for i := uint(0); i < cdll.length; i++ {
		cur = cur.next
	}
	return cur
}

func (cdll *CircularDoublyLinkedList) Append(v interface{}) {
	if cdll.head.next == nil {
		newNode := NewNodeCD(v)
		newNode.next = newNode
		newNode.prev = newNode
		cdll.head.next = newNode
	} else {
		head := cdll.GetHead()
		tail := cdll.GetTail()
		newNode := NewNodeCD(v)

		tail.next = newNode
		newNode.prev = tail

		newNode.next = head.next
		head.next.prev = newNode
	}
	cdll.length++
}

func (cdll *CircularDoublyLinkedList) FindByIndex(idx uint) (*NodeCD, error) {
	if idx >= cdll.length {
		return nil, errors.New("out of range")
	}
	cur := cdll.head.next
	for i := uint(0); i < idx; i++ {
		cur = cur.next
	}
	return cur, nil
}

func (cdll *CircularDoublyLinkedList) InsertAfter(idx uint, v interface{}) (bool, error) {
	if idx >= cdll.length {
		return false, errors.New("out of range")
	}
	cur, _ := cdll.FindByIndex(idx)
	newNode := NewNodeCD(v)

	newNode.next = cur.next
	cur.next.prev = newNode
	cur.next = newNode
	newNode.prev = cur

	cdll.length++
	return true, nil
}

func (cdll *CircularDoublyLinkedList) InsertBefore(idx uint, v interface{}) (bool, error) {
	if idx >= cdll.length {
		return false, errors.New("out of range")
	}
	cur, _ := cdll.FindByIndex(idx)
	newNode := NewNodeCD(v)

	cur.prev.next = newNode
	newNode.prev = cur.prev
	newNode.next = cur
	cur.prev = newNode

	cdll.length++
	return true, nil

}

func (cdll *CircularDoublyLinkedList) Remove(idx uint) interface{} {
	if idx >= cdll.length {
		return nil
	}
	cur, _ := cdll.FindByIndex(idx)

	v := cur.val
	cur.prev.next = cur.next
	cur.next.prev = cur.prev
	cur = nil

	cdll.length--
	return v
}

func (cdll *CircularDoublyLinkedList) Print() {
	cur := cdll.head.next
	format := ""
	for i := uint(0); i < cdll.length+1; i++ {
		format += fmt.Sprintf("%+v", cur.val)
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}
