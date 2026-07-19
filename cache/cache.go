package cache

import (
	"sync"
	"time"
)

type entry[V any] struct {
	expirable bool
	val       V
	expiresAt time.Time
}

type Cache[K comparable, V any] struct {
	data map[K]entry[V]
	mu   sync.RWMutex
}

func (c *Cache[K, V]) Set(key K, data V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = entry[V]{
		val:       data,
		expirable: false,
		expiresAt: time.Time{},
	}
}

func (c *Cache[K, V]) SetWithTTL(key K, data V, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = entry[V]{
		val:       data,
		expirable: true,
		expiresAt: time.Now().Add(duration),
	}
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	var zero V
	if val.expirable && time.Now().After(val.expiresAt) {
		return zero, false
	}
	return val.val, ok
}

func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		mu:   sync.RWMutex{},
		data: make(map[K]entry[V]),
	}
}
