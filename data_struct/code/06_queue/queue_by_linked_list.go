package _6_queue

import (
	"errors"
	"fmt"
)

type Node struct {
	val  interface{}
	next *Node
}

type LinkedListQueue struct {
	head     *Node
	tail     *Node
	size     int
	capacity int
}

func NewLinkedListQueue(capacity int) *LinkedListQueue {
	return &LinkedListQueue{
		head:     nil,
		tail:     nil,
		size:     0,
		capacity: capacity,
	}
}

func (q *LinkedListQueue) Enqueue(v interface{}) error {
	if q.size == q.capacity {
		return errors.New("queue is full")
	}
	node := &Node{v, nil}
	if q.IsEmpty() {
		q.tail = node
		q.head = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.size++
	return nil
}

func (q *LinkedListQueue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	v := q.head.val
	q.head = q.head.next
	q.size--
	return v, nil
}

func (q *LinkedListQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LinkedListQueue) Print() error {
	if q.IsEmpty() {
		return errors.New("queue is empty")
	}
	for cur := q.head; cur != nil; cur = cur.next {
		fmt.Printf("%v → ", cur.val)
	}
	return nil
}
