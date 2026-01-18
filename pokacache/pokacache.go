package pokecache

import (
	"sync"
	"time"
)

// Cache stores byte slices by key with a simple time-based eviction strategy.
type Cache struct {
	cache map[string]cacheEntry
	// Mutex guards concurrent access to the map.
	mux *sync.Mutex
}

// cacheEntry records the stored data and the time it was added.
type cacheEntry struct {
	data      []byte
	createdAt time.Time
}

// NewCache builds a Cache and starts a background reaper goroutine.
func NewCache( interval time.Duration) *Cache {
	c := &Cache{
		cache: make(map[string]cacheEntry),
		mux: &sync.Mutex{},
	}
	// "go" starts loopReap in a new goroutine.
	go c.loopReap(interval)

	return c
}


// Get returns the cached data and a bool that indicates whether it exists.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, exists := c.cache[key]
	if !exists {
		return nil, false
	}
	return entry.data, true
}

// Add inserts or replaces a cache entry.
func (c *Cache) Add(key string, data []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		data:      data,
		createdAt: time.Now(),
	}
}
// Delete removes an entry by key.
func (c *Cache) Delete(key string) {
	delete(c.cache, key)
}
// Clear returns the cached data if present (note: it does not delete it).
func (c *Cache) Clear(key string) (data []byte) {
	entry, exists := c.cache[key]
	if !exists {
		return nil
	}
	return entry.data
}

// loopReap runs forever, periodically calling reap using a ticker channel.
func (c *Cache) loopReap(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C{
		c.reap(interval)
	}
}


// reap deletes entries older than the interval.
func(c *Cache) reap(interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	expiryTime:= time.Now().UTC().Add(-interval)
	for key, value := range c.cache {
		if value.createdAt.Before(expiryTime) {
			delete(c.cache, key)
		}
	}
}
