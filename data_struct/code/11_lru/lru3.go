package _11_lru

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
)

func lru3() {
	cache, _ := lru.New(5)
	cache.Add("key1", "value1")
	cache.Add("key2", "value2")
	value, ok := cache.Get("key1")
	if ok {
		fmt.Println(value)
	}
	cache.Remove("key2")
}
