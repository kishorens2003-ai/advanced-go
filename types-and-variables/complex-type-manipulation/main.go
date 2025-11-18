package main

import (
	"sync"
	"time"
)

type CacheItem struct {
	Value      any
	Expiration time.Time
}

func (c CacheItem) IsExpired() bool {
	return time.Now().After(c.Expiration)
}

type SafeCache struct {
	mu    sync.RWMutex
	items map[string]CacheItem
	stop  chan struct{}
}

func NewSafeCache() *SafeCache {
	cache := &SafeCache{
		items: make(map[string]CacheItem),
		stop:  make(chan struct{}),
	}

	// Start cleanup goroutine
	go cache.cleanup()
	return cache
}

func (c *SafeCache) Set(key string, value any, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = CacheItem{
		Value:      value,
		Expiration: time.Now().Add(ttl),
	}
}

func (c *SafeCache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.items[key]
	if !exists || item.IsExpired() {
		return nil, false
	}

	return item.Value, true
}

func (c *SafeCache) GetTyped(key string, target any) bool {
	value, exists := c.Get(key)
	if !exists {
		return false
	}

	// Type assertion with reflection would be more robust
	switch v := target.(type) {
	case *string:
		if str, ok := value.(string); ok {
			*v = str
			return true
		}
	case *int:
		if num, ok := value.(int); ok {
			*v = num
			return true
		}
	case *map[string]any:
		if m, ok := value.(map[string]any); ok {
			*v = m
			return true
		}
	}

	return false
}

func (c *SafeCache) cleanup() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			for key, item := range c.items {
				if item.IsExpired() {
					delete(c.items, key)
				}
			}
			c.mu.Unlock()
		case <-c.stop:
			return
		}
	}
}

func (c *SafeCache) Close() {
	close(c.stop)
}
