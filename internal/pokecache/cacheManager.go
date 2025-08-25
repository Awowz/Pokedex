package pokecache

import (
	"sync"
	"time"
)

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheMap[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c Cache) Get(key string) ([]byte, bool) {
	data, found := c.cacheMap[key]
	return data.val, found
}

func NewCache(interval time.Duration) Cache {
	myCache := Cache{mu: sync.Mutex{}, cacheMap: make(map[string]cacheEntry), interval: interval}
	go myCache.reapLoop()
	return myCache
}

func (c Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for {
		current_time := <-ticker.C
		for key := range c.cacheMap {
			if float64(c.cacheMap[key].createdAt.Second())+c.interval.Seconds() <= float64(current_time.Second()) {
				c.mu.Lock()
				delete(c.cacheMap, key)
				c.mu.Unlock()
			}
		}
	}
}
