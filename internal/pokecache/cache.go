package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mu    sync.RWMutex
}
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.entry[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entry: make(map[string]cacheEntry),
	}
	go c.reapLoop(time.Duration(interval))
	return c
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			cache.mu.Lock()
			for item, data := range cache.entry {
				if time.Since(data.createdAt) > interval {
					delete(cache.entry, item)
				}
			}
			cache.mu.Unlock()
		}
	}
}
