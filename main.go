package main

import (
	"encoding/json"
	"fmt"
)

type Cache[K comparable, V any] struct {
	data map[K]V
}

func (c *Cache[K, V]) Set(key K, data V) {
	c.data[key] = data
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	val, ok := c.data[key]
	return val, ok
}

func (c *Cache[K, V]) Delete(key K) {
	delete(c.data, key)
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		data: make(map[K]V),
	}
}

func main() {
	cache := NewCache[string, []byte]()
	key := "name"

	data, _ := json.Marshal("aslamcodes")
	cache.Set(key, data)

	if _, ok := cache.Get(key); !ok {
		fmt.Println("not found")
	}

}
