package pokecache

import (
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

//reminder to add mutex to struct
