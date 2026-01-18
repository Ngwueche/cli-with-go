package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux *sync.Mutex
}

type cacheEntry struct {
	data      []byte
	createdAt time.Time
}

func NewCache( interval time.Duration) *Cache {
	c := &Cache{
		cache: make(map[string]cacheEntry),
		mux: &sync.Mutex{},
	}
	go c.loopReap(interval)

	return c
}


func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, exists := c.cache[key]
	if !exists {
		return nil, false
	}
	return entry.data, true
}

func (c *Cache) Add(key string, data []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		data:      data,
		createdAt: time.Now(),
	}
}
func (c *Cache) Delete(key string) {
	delete(c.cache, key)
}
func (c *Cache) Clear(key string) (data []byte) {
	entry, exists := c.cache[key]
	if !exists {
		return nil
	}
	return entry.data
}

func (c *Cache) loopReap(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C{
		c.reap(interval)
	}
}


func(c *Cache) reap(interval time.Duration) {
	// implement cache expiration logic if needed
	c.mux.Lock()
	defer c.mux.Unlock()
	expiryTime:= time.Now().UTC().Add(-interval)
	for key, value := range c.cache {
		if value.createdAt.Before(expiryTime) {
			delete(c.cache, key)
		}
	}
}
