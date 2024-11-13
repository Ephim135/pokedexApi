package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu sync.Mutex
	cacheEntries map[string]cacheEntry
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		mu: var mu sync.Mutex
		cacheEntries: make(map[string]cacheEntry),
		interval: interval,
	}
	go cache.startCleanup()
	return cache
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheEntries[key].val = val
	c.cacheEntries[key].createdAt = time.Now()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.cacheEntries[key]
}

func (c Cache) startCleanup() (){
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		defer c.mu.Unlock()
		// Loop over Entries in Cache
		for key, _ := range c.cacheEntries {
			duration := Since(c.cacheEntries[key].createdAt)
			if duration > c.interval {
				delete(c.cacheEntries, key)
			}
	}
}
}