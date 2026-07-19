package cache

import "sync"

type Cache[K comparable, V any] struct {
	data map[K]V
	mu   sync.RWMutex
}

func (c *Cache[K, V]) Set(key K, data V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = data
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	return val, ok
}

func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		mu:   sync.RWMutex{},
		data: make(map[K]V),
	}
}
