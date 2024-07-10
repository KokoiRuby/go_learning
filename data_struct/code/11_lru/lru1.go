package _11_lru

import "fmt"

type LRUNode1 struct {
	key   string
	value interface{}
	next  *LRUNode1
	prev  *LRUNode1
}

type LRUCache1 struct {
	m          map[string]*LRUNode1
	head, tail *LRUNode1
	capacity   int
	used       int
}

func NewLRUCache1(capacity int) *LRUCache1 {
	return &LRUCache1{
		m:        make(map[string]*LRUNode1, capacity),
		head:     &LRUNode1{},
		tail:     &LRUNode1{},
		capacity: capacity,
		used:     0,
	}
}

func (c *LRUCache1) Get(key string) interface{} {
	if c.isEmpty() {
		return nil
	}
	// search in map
	node, ok := c.m[key]
	if ok {
		// gotcha
		c.moveToHead(node)
		return node.value
	} else {
		return nil
	}

}

func (c *LRUCache1) Put(key string, value interface{}) {
	// hit
	if node, ok := c.searchNode(key); ok {
		c.moveToHead(node)
		return
		// miss
	} else {
		if c.used == c.capacity {
			// full
			c.delTail()
			c.addToHead(key, value)
			return

		} else {
			// not full
			c.addToHead(key, value)
			return
		}
	}
}

func (c *LRUCache1) Remove(key string) {
	if node, ok := c.searchNode(key); ok {
		node.prev.next = node.next
		node.next.prev = node.prev
		node = nil
		delete(c.m, key)
	} else {
		return
	}
}

func (c *LRUCache1) moveToHead(node *LRUNode1) {
	if node == c.head {
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev

	c.tail.next = node
	node.prev = c.tail
	node.next = c.head
	c.head.prev = node
	c.head = node
	return
}

func (c *LRUCache1) addToHead(k string, v interface{}) {
	newNode := &LRUNode1{
		key:   k,
		value: v,
	}

	// cache is empty
	if c.isEmpty() {
		c.head, c.tail = newNode, newNode
		c.used++
		return
	}

	c.tail.next = newNode
	newNode.prev = c.tail
	newNode.next = c.head
	c.head.prev = newNode
	c.head = newNode

	// save to map
	c.m[k] = newNode
	c.used++
	return
}

func (c *LRUCache1) delTail() {
	if c.isEmpty() {
		return
	}
	c.tail.prev.next = c.head
	c.head.prev = c.tail.prev
	c.tail = c.tail.prev
	c.used--
	return
}

func (c *LRUCache1) searchNode(key string) (*LRUNode1, bool) {
	node, ok := c.m[key]
	return node, ok
}

func (c *LRUCache1) isEmpty() bool {
	return c.used == 0
}

func (c *LRUCache1) isFull() bool {
	return c.used == c.capacity
}

func (c *LRUCache1) Print() {
	count := 0
	for cur := c.head; count <= c.capacity; cur = cur.next {
		fmt.Printf("%v -> ", cur.key)
		count++
	}
}
