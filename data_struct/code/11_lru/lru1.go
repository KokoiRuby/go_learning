package _11_lru

type _LRUNode1 struct {
	key   string
	value interface{}
	next  *LRUNode
	prev  *LRUNode
}

type LRUCache1 struct {
	m          []map[string]*_LRUNode1
	head, tail *_LRUNode1
	used       int
}

func New_LRUCache(capacity int) *LRUCache1 {
	return &LRUCache1{
		m:    make([]map[string]*_LRUNode1, capacity),
		head: &_LRUNode1{},
		tail: &_LRUNode1{},
		used: 0,
	}
}
