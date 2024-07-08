package _6_queue

import (
	"errors"
	"fmt"
)

type CircularQueue struct {
	q        []interface{}
	head     int
	tail     int
	capacity int
}

func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		q:        make([]interface{}, capacity),
		head:     0,
		tail:     0,
		capacity: capacity,
	}
}

func (q *CircularQueue) IsEmpty() bool {
	return q.head == q.tail
}

func (q *CircularQueue) IsFull() bool {
	return (q.tail+1)%q.capacity == q.head
}

func (q *CircularQueue) Enqueue(v interface{}) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}
	q.q[q.tail] = v
	q.tail = (q.tail + 1) % q.capacity
	return nil
}

func (q *CircularQueue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	v := q.q[q.head]
	q.head = (q.head + 1) % q.capacity
	return v, nil
}

func (q *CircularQueue) Print() {
	for _, v := range q.q {
		fmt.Printf("%v →", v)
	}
}
