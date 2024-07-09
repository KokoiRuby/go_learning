package _11_lru

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

func New_LRUCache(capacity int) *LRUCache1 {
	return &LRUCache1{
		m:        make(map[string]*LRUNode1, capacity),
		head:     &LRUNode1{},
		tail:     &LRUNode1{},
		capacity: capacity,
		used:     0,
	}
}

func (c *LRUCache1) Get(key string) interface{} {
	if c.used == 0 {
		return nil
	}
	node, ok := c.m[key]
	if ok {
		c.moveToHead(node)
		return node.value
	} else {
		return nil
	}

}

func (c *LRUCache1) Put(key string, value interface{}) {
	// hit
	if node := c.searchNode(key); node != nil {
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
	return
}

func (c *LRUCache1) moveToHead(node *LRUNode1) {
	return
}

func (c *LRUCache1) addToHead(key string, value interface{}) {
	return
}

func (c *LRUCache1) delTail() {
	return
}

func (c *LRUCache1) searchNode(key string) *LRUNode1 {
	return nil
}
