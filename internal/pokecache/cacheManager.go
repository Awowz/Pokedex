package pokecache

import (
	"fmt"
	"sync"
	"time"
)

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheMap[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	data, found := c.cacheMap[key]
	return data.val, found
}

func NewCache(interval time.Duration) Cache {
	myCache := Cache{mu: &sync.Mutex{}, cacheMap: make(map[string]cacheEntry), interval: interval}
	go myCache.reapLoop()
	return myCache
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for {
		current_time := <-ticker.C
		for key := range c.cacheMap {
			dataEntryEndTime := c.cacheMap[key].createdAt.Add(c.interval)
			if current_time.Equal(dataEntryEndTime) || current_time.After(dataEntryEndTime) {
				fmt.Printf("reaping...")
				c.mu.Lock()
				delete(c.cacheMap, key)
				c.mu.Unlock()
			}
		}
	}
}
