package _4_linked_list

import (
	"errors"
	"fmt"
)

type Node struct {
	next *Node
	val  interface{}
}

type SinglyLinkedList struct {
	head   *Node
	length uint
}

func NewNode(v interface{}) *Node {
	return &Node{val: v}
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		head:   NewNode(nil),
		length: 0,
	}
}

func (sll *SinglyLinkedList) GetValue(node *Node) interface{} {
	return node.val
}

func (sll *SinglyLinkedList) GetHead() *Node {
	return sll.head
}

func (sll *SinglyLinkedList) GetTail() *Node {
	cur := sll.head
	for cur.next != nil {
		cur = cur.next
	}
	return cur
}

func (sll *SinglyLinkedList) Append(v interface{}) {
	if sll.head.next == nil {
		sll.head.next = NewNode(v)
	} else {
		tail := sll.GetTail()
		tail.next = NewNode(v)
	}
	sll.length++
}

func (sll *SinglyLinkedList) FindByIndex(idx uint) (prev *Node, curr *Node, err error) {
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

func (sll *SinglyLinkedList) InsertAfter(idx uint, v interface{}) (bool, error) {
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

func (sll *SinglyLinkedList) InsertBefore(idx uint, v interface{}) (bool, error) {
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

func (sll *SinglyLinkedList) Remove(idx uint) interface{} {
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

func (sll *SinglyLinkedList) Print() {
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

/*
	Extra
*/

func (sll *SinglyLinkedList) Reverse() {
	if sll.head == nil || sll.head.next == nil || sll.head.next.next == nil {

		return
	}

	var pre *Node
	cur := sll.head.next
	for cur != nil {
		next := cur.next
		cur.next = pre
		// step in
		pre = cur
		cur = next
	}
	sll.head.next = pre
}

func (sll *SinglyLinkedList) HasCycle() bool {
	if sll.head == nil {
		return false
	}
	slow := sll.head
	fast := sll.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

func MergeSortedList(sll1, sll2 *SinglyLinkedList) *SinglyLinkedList {
	if sll1.head == nil || sll1.head.next == nil || sll1.head.next.next == nil {
		return sll2
	}
	if sll2.head == nil || sll2.head.next == nil || sll2.head.next.next == nil {
		return sll1
	}

	msll := NewSinglyLinkedList()
	cur := msll.head
	cur1 := sll1.head.next
	cur2 := sll2.head.next
	for cur1 != nil && cur2 != nil {
		if cur1.val.(int) > cur2.val.(int) {
			cur.next = cur2
			cur2 = cur2.next
		} else {
			cur.next = cur1
			cur1 = cur1.next
		}
		cur = cur.next
	}

	// deal with left
	if cur1 != nil {
		cur.next = cur1
	} else if cur2 != nil {
		cur.next = cur2
	}

	return msll
}

func (sll *SinglyLinkedList) FindMiddleNode() *Node {
	if sll.head == nil || sll.head.next == nil {
		return nil
	}
	slow, fast := sll.head, sll.head
	for slow != nil && fast != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

func (sll *SinglyLinkedList) RemoveBottomN(n int) {
	if n <= 0 || sll.head == nil || sll.head.next == nil {
		return
	}
	fast := sll.head
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.next
	}
	if fast != nil {
		slow := sll.head
		for fast.next != nil {
			slow = slow.next
			fast = fast.next
		}
		slow.next = slow.next.next
		slow = nil
	}

}
