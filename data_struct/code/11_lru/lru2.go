package _11_lru

import "hash/fnv"

const CAPACITY = 100

type LRUNode struct {
	key   string
	value interface{}
	next  *LRUNode
	prev  *LRUNode
	hnext *LRUNode
}

type LRUCache struct {
	slots []LRUNode // singly

	head, tail *LRUNode // doubly

	used int
}

func NewLRUCache() *LRUCache {
	return &LRUCache{
		slots: make([]LRUNode, CAPACITY),
		head:  nil,
		tail:  nil,
		used:  0,
	}
}

func (c *LRUCache) Get(key string) interface{} {
	if c.used == 0 {
		return nil
	}
	if tmp := c.searchNode(key); tmp != nil {
		c.moveToHead(tmp)
		return tmp.value
	}
	return nil
}

func (c *LRUCache) Put(key string, value interface{}) {
	// hit
	if tmp := c.searchNode(key); tmp != nil {
		c.moveToHead(tmp)
	}
	// miss, full
	if c.used == CAPACITY {
		c.delTail()
		c.addToHead(key, value)
	}
	// miss, not full
	c.addToHead(key, value)
}

func (c *LRUCache) addToHead(k string, v interface{}) {
	newNode := &LRUNode{
		key:   k,
		value: v,
	}

	// cache is empty
	if c.tail == nil {
		c.head, c.tail = newNode, newNode
		return
	}

	// singly
	tmpSlot := &c.slots[simpleHash(k)]
	newNode.hnext = tmpSlot.hnext
	tmpSlot.hnext = newNode
	c.used++

	// doubly
	c.tail.next = newNode
	newNode.prev = c.tail
	c.tail = newNode

}

// miss & full
func (c *LRUCache) delTail() {
	// cache is empty
	if c.tail == nil {
		return
	}
	tailSlot := &c.slots[simpleHash(c.tail.key)]
	if tailSlot.hnext != nil {
		// singly
		tailSlot = tailSlot.hnext
		// doubly
		tailSlot.prev.next = tailSlot.next
		tailSlot.next.prev = tailSlot.prev
		tailSlot = nil
		c.used--
	} else {
		// doubly
		tailSlot.prev.next = tailSlot.next
		tailSlot.next.prev = tailSlot.prev
		tailSlot = nil
		c.used--
	}
}

// hit
func (c *LRUCache) moveToHead(node *LRUNode) {
	if node == c.head {
		return
	}
	// doubly
	node.prev.next = node.next
	node.next.prev = node.prev

	c.tail.next = node
	node.prev = c.tail
	node.next = c.head
	c.head.prev = node
	c.head = node
}

func (c *LRUCache) searchNode(key string) *LRUNode {
	return nil
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
