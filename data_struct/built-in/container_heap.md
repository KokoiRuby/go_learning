Package heap provides heap operations for any type that implements **heap.Interface**. 

Each node is the **minimum-valued** node in its subtree. root at index 0.

```go
// interface to impl for a type

```



```go
package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

// Implementing heap.Interface methods for IntHeap
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds an element to the heap.
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the smallest element from the heap.
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	// Creating a new IntHeap
	h := &IntHeap{2, 1, 5}
	heap.Init(h) // Initializes the heap in-place

	// Pushing new elements onto the heap
	heap.Push(h, 3)
	fmt.Printf("After pushing 3: %v\n", *h)

	// Pushing another element onto the heap
	heap.Push(h, 4)
	fmt.Printf("After pushing 4: %v\n", *h)

	// Popping the smallest element from the heap
	fmt.Printf("Popped: %v\n", heap.Pop(h))
	fmt.Printf("Heap after popping: %v\n", *h)

	// Popping the smallest element from the heap
	fmt.Printf("Popped: %v\n", heap.Pop(h))
	fmt.Printf("Heap after popping: %v\n", *h)
    
    // Fix: Assume we change the value of an element
	(*h)[0] = 6
	heap.Fix(h, 0)
	fmt.Printf("After fixing: %v\n", *h)

	// Remove: Removing the element at index 2
	fmt.Printf("Removed: %v\n", heap.Remove(h, 2))
	fmt.Printf("Heap after removing: %v\n", *h)
}

```

### Priority Queue

A heap is a common way to implement a **priority queue**. To build a priority queue, implement the Heap interface with the **(negative) priority** as the ordering for the Less method, so Push adds items while **Pop removes the highest-priority item** from the queue.

```go
// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func main() {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
}

```

