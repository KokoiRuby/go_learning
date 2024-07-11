package _11_lru

import (
	"fmt"
	"hash/fnv"
)

const CAPACITY = 100

type LRUNode struct {
	key   string
	value interface{}
	next  *LRUNode
	prev  *LRUNode
	hnext *LRUNode
}

type LRUCache struct {
	slots      []*LRUNode // singly
	head, tail *LRUNode   // doubly
	capactiy   int
	used       int
}

func NewLRUCache(capactiy int) *LRUCache {
	return &LRUCache{
		slots:    make([]*LRUNode, CAPACITY),
		head:     nil,
		tail:     nil,
		capactiy: capactiy,
		used:     0,
	}
}

func (c *LRUCache) Get(key string) interface{} {
	if c.isEmpty() {
		return nil
	}
	if node := c.searchNode(key); node != nil {
		c.moveToHead(node)
		return node.value
	} else {
		return nil
	}
}

func (c *LRUCache) Put(key string, value interface{}) {
	// hit
	if node := c.searchNode(key); node != nil {
		c.moveToHead(node)
		return
	} else {
		if c.isFull() {
			// miss + full
			c.delTail()
			c.addToHead(key, value)
			return

		} else {
			// miss + not full
			c.addToHead(key, value)
			return
		}
	}
}

func (c *LRUCache) Remove(key string) {
	if node := c.searchNode(key); node != nil {
		// doubly
		node.prev.next = node.next
		node.next.prev = node.prev

		// singly
		var prev *LRUNode
		slot := c.slots[simpleHash(key)]
		for node := slot; node != nil; {
			// gotcha
			if node.key == key {
				// slot
				if prev == nil {
					c.slots[simpleHash(key)] = node.hnext
				} else {
					prev.hnext = node.hnext
				}
				node = nil
				c.used--
				return
			}
			// step forward
			prev = node
			node = node.hnext
		}
	} else {
		return
	}
}

func (c *LRUCache) addToHead(k string, v interface{}) {
	newNode := &LRUNode{
		key:   k,
		value: v,
	}

	// cache is empty
	if c.isEmpty() {
		c.head, c.tail = newNode, newNode
		c.used++
		return
	}

	// singly
	slot := c.slots[simpleHash(k)]
	if slot == nil {
		c.slots[simpleHash(k)] = newNode
	} else {
		newNode.hnext = slot
		slot = newNode
	}

	// doubly
	c.tail.next = newNode
	newNode.prev = c.tail
	newNode.next = c.head
	c.head.prev = newNode
	c.head = newNode

	c.used++

}

func (c *LRUCache) delTail() {
	if c.isEmpty() {
		return
	}
	// singly
	var prev *LRUNode
	tailSlot := c.slots[simpleHash(c.tail.key)]
	for node := tailSlot; node != nil; {
		// gotcha
		if node.key == c.tail.key {
			if node.hnext == nil {
				break
			} else {
				prev.hnext = node.hnext
				break
			}
		}
		// step forward
		prev = node
		node = node.hnext
	}

	// doubly
	c.tail.prev.next = c.head
	c.head.prev = c.tail.prev

	c.used--
}

func (c *LRUCache) moveToHead(node *LRUNode) {
	if node == c.head {
		return
	}
	// doubly only
	node.prev.next = node.next
	node.next.prev = node.prev

	c.tail.next = node
	node.prev = c.tail
	node.next = c.head
	c.head.prev = node
	c.head = node
}

func (c *LRUCache) searchNode(key string) *LRUNode {
	if c.isEmpty() {
		return nil
	}
	slot := c.slots[simpleHash(key)]
	if slot == nil {
		return nil
	} else {
		for node := slot; node != nil; node = node.hnext {
			if node.key == key {
				return node
			}
		}
		return nil
	}

}

func (c *LRUCache) isEmpty() bool {
	return c.used == 0
}

func (c *LRUCache) isFull() bool {
	return c.used == c.capactiy
}

func simpleHash(data string) int {
	hash := fnv.New32a()
	_, err := hash.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	hashValue := hash.Sum32()
	return int(hashValue) % CAPACITY
}

func (c *LRUCache) Print() {
	count := 0
	for cur := c.head; count <= c.capactiy; cur = cur.next {
		fmt.Printf("%v -> ", cur.key)
		count++
	}
}
