package pokecache

import (
	"sync"
	"time"
)


func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	c.CacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val: value,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	value, ok := c.CacheMap[key]
	c.mu.Unlock()
	if !ok {
		return nil, false
	}
	return value.val, true

}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
		t := <-ticker.C

		c.mu.Lock()
		for k, v := range c.CacheMap {
			if t.Sub(v.createdAt) > c.interval {
				delete(c.CacheMap, k)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		CacheMap: make(map[string]cacheEntry),
		interval: interval,
		mu: &sync.Mutex{},
	}

	go cache.reapLoop()

	return cache
}