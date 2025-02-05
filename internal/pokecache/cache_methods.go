package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {

	c := &Cache{
		interval:  interval,
		mu:        sync.Mutex{},
		cacheData: map[string]cacheEntry{},
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cacheEntry, exists := c.cacheData[key]
	if !exists {
		return []byte{}, exists
	}
	return cacheEntry.val, exists
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheData[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		t := <-ticker.C
		c.mu.Lock()
		for key, cacheEntry := range c.cacheData {
			if t.Sub(cacheEntry.createdAt) >= c.interval {

				delete(c.cacheData, key)
			}
		}
		c.mu.Unlock()
	}
}
