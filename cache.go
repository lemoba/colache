package main

import (
	"log"
	"sync"
	"time"
)

var (
	cache = make(map[string]*CacheTable)
	mutex sync.RWMutex
)

// Cache returns the existing cache table with given name or creates a new one
// if the table does not exist yet.
func Cache(table string) *CacheTable {
	mutex.RLock()
	t, ok := cache[table]
	mutex.RUnlock()

	if !ok {
		mutex.Lock()
		t, ok = cache[table]
		// Double check whether the table exists or not.
		if !ok {
			t = &CacheTable{
				name:  table,
				items: make(map[interface{}]*CacheItem),
			}
			cache[table] = t
		}
		mutex.Unlock()
	}

	return t
}

type CacheTable struct {
	sync.RWMutex

	// The table's name.
	name string

	// All Cached itmes.
	items map[interface{}]*CacheItem

	// Timer responsible for triggering cleanup.
	cleanupTimer *time.Timer
	// Current timer duration.
	cleanupInterval time.Duration

	// The logger used for this table.
	logger *log.Logger

	// Callback method triggered when trying to load a non-existing key.
	loadData func(key interface{}, args ...interface{}) *CacheItem
	// Callback method triggered when adding a new item to the cache.
	addedItem []func(item *CacheItem)
	// Callback method triggered before deleting an item from the cache.
	aboutToDeleteItem []func(item *CacheItem)
}
