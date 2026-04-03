// Package pokecache...
package pokecache

import (
	"sync"
	"time"
)

var LocationCache = NewCache(5 * time.Second)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	CacheEntries map[string]cacheEntry
	Mu           sync.Mutex
	Interval     time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		CacheEntries: make(map[string]cacheEntry),
		Interval:     interval,
	}

	go cache.reapLoop()

	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.Mu.Lock()
	defer cache.Mu.Unlock()

	cache.CacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.Mu.Lock()
	defer cache.Mu.Unlock()
	if entry, exists := cache.CacheEntries[key]; exists {
		return entry.val, true
	}

	return []byte{}, false
}

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.Interval)

	defer ticker.Stop()

	for range ticker.C {
		cache.Mu.Lock()

		for key, entry := range cache.CacheEntries {
			if time.Since(entry.createdAt) > cache.Interval {
				delete(cache.CacheEntries, key)
			}
		}

		cache.Mu.Unlock()
	}
}
