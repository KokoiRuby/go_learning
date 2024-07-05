package _4_linked_list

import (
	"errors"
	"fmt"
)

type NodeD struct {
	prev *NodeD
	next *NodeD
	val  interface{}
}

type DoublyLinkedList struct {
	head   *NodeD
	length uint
}

func NewNodeD(v interface{}) *NodeD {
	return &NodeD{
		val: v,
	}
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head:   NewNodeD(nil),
		length: 0,
	}
}

func (dll *DoublyLinkedList) GetHead() *NodeD {
	return dll.head
}

func (dll *DoublyLinkedList) GetTail() *NodeD {
	cur := dll.head
	for cur.next != nil {
		cur = cur.next
	}
	return cur
}

func (dll *DoublyLinkedList) Append(v interface{}) {
	if dll.head.next == nil {
		dll.head.next = NewNodeD(v)
	} else {
		tail := dll.GetTail()
		newNode := NewNodeD(v)
		tail.next = newNode
		newNode.prev = tail
	}
	dll.length++
}

func (dll *DoublyLinkedList) FindByIndex(idx uint) (*NodeD, error) {
	if idx >= dll.length {
		return nil, errors.New("out of range")
	}
	cur := dll.head.next
	for i := uint(0); i < idx; i++ {
		cur = cur.next
	}
	return cur, nil
}

func (dll *DoublyLinkedList) InsertAfter(idx uint, v interface{}) (bool, error) {
	if idx >= dll.length {
		return false, errors.New("out of range")
	}
	cur, _ := dll.FindByIndex(idx)
	newNode := NewNodeD(v)

	newNode.next = cur.next
	cur.next.prev = newNode
	cur.next = newNode
	newNode.prev = cur

	dll.length++
	return true, nil
}

func (dll *DoublyLinkedList) InsertBefore(idx uint, v interface{}) (bool, error) {
	if idx >= dll.length {
		return false, errors.New("out of range")
	}
	cur, _ := dll.FindByIndex(idx)
	newNode := NewNodeD(v)

	cur.prev.next = newNode
	newNode.prev = cur.prev
	newNode.next = cur
	cur.prev = newNode

	dll.length++
	return true, nil

}

func (dll *DoublyLinkedList) Remove(idx uint) interface{} {
	if idx >= dll.length {
		return nil
	}
	cur, _ := dll.FindByIndex(idx)

	v := cur.val
	cur.prev.next = cur.next
	cur.next.prev = cur.prev
	cur = nil

	dll.length--
	return v
}

func (dll *DoublyLinkedList) Print() {
	cur := dll.head.next
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
