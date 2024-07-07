package _6_queue

import (
	"errors"
	"fmt"
)

type ArrayQueue struct {
	q    []interface{}
	head int
	tail int
	size int
}

func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		q:    make([]interface{}, capacity),
		head: 0,
		tail: 0,
	}
}

func (q *ArrayQueue) Enqueue(v interface{}) error {
	// reach the end of array → migrate
	if q.tail == cap(q.q) && q.head != 0 {
		for i := q.head; i < q.tail; i++ {
			q.q[i-q.head] = q.q[i]
		}
		q.tail = q.tail - q.head
		q.head = 0

	}
	if q.size == cap(q.q) {
		return errors.New("queue is full")
	}
	q.q[q.tail] = v
	q.tail++
	q.size++
	return nil
}

func (q *ArrayQueue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	v := q.q[q.head]
	q.head++
	q.size--
	return v, nil
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.size == 0
}

func (q ArrayQueue) Print() {
	for i := 0; i < cap(q.q); i++ {
		fmt.Println(q.q[i])
	}
}
