package _12_heap

import "fmt"

type Heap struct {
	a        []interface{}
	capacity int
	size     int
}

func NewHeap(capacity int) *Heap {
	return &Heap{
		a:        make([]interface{}, capacity),
		capacity: capacity,
		size:     0,
	}
}

func (h *Heap) Insert(v interface{}) {
	if h.isFull() {
		return
	}

	// append
	h.a[h.size] = v

	// compare with parent from bottom to top
	child := h.size
	parent := (child - 1) / 2
	for parent >= 0 && h.isLess(h.a[parent], h.a[child]) {
		// swap if parent is less than child
		h.swap(h.a, parent, child)
		child = parent
		parent = (child - 1) / 2
	}
	h.size++
}

func (h *Heap) Pop() interface{} {
	if h.isEmpty() {
		return nil
	}

	// bottom to top
	ret := h.a[0]
	h.a[0] = h.a[h.size-1]
	h.a[h.size-1] = nil
	h.size--

	// heapify
	h.heapify(0)
	return ret

}

func (h *Heap) heapify(i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// Compare with left child
	if left < h.size && h.a[left].(int) > h.a[largest].(int) {
		largest = left
	}

	// Compare with right child
	if right < h.size && h.a[right].(int) > h.a[largest].(int) {
		largest = right
	}

	// If the smallest is not the current node, swap and continue heapifying
	if largest != i {
		h.swap(h.a, i, largest)
		h.heapify(largest)
	}
}

func (h *Heap) isEmpty() bool {
	return h.size == 0
}

func (h *Heap) isFull() bool {
	return h.size == h.capacity
}

// func (h *Heap) isGreater(a, b interface{}) bool {
// 	aInt := a.(int)
// 	bInt := b.(int)

// 	if aInt > bInt {
// 		return true
// 	} else {
// 		return false
// 	}
// }

func (h *Heap) isLess(a, b interface{}) bool {
	aInt := a.(int)
	bInt := b.(int)

	if aInt < bInt {
		return true
	} else {
		return false
	}
}

func (h *Heap) swap(a []interface{}, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (h *Heap) Print() {
	for i := 0; i < len(h.a); i++ {
		fmt.Printf("%v ", h.a[i])
	}
}
