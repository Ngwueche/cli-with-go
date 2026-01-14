package pokacache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	data      []byte
	createdAt time.Time
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]cacheEntry),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, exists := c.cache[key]
	if !exists {
		return nil, false
	}
	return entry.data, true
}

func (c *Cache) Add(key string, data []byte) {
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
