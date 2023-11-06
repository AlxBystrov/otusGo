package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}
type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	// when the key is in items, just move to front and update the value
	if _, ok := c.items[key]; ok {
		c.items[key].Value = cacheItem{
			key:   key,
			value: value,
		}
		c.queue.MoveToFront(c.items[key])
		return true
	}
	// when the key is not in cache, check the capacity, remove back() when the cache is full
	if c.queue.Len() == c.capacity {
		item := c.queue.Back().Value.(cacheItem)
		delete(c.items, item.key)
		c.queue.Remove(c.queue.Back())
	}
	item := cacheItem{
		key:   key,
		value: value,
	}
	c.items[key] = c.queue.PushFront(item)
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// return value if key in items
	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)
		return item.Value.(cacheItem).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
