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
	keys     map[*ListItem]Key
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	// when the key is in items, just move to front and update the value
	if _, ok := c.items[key]; ok {
		c.items[key].Value = value
		c.queue.MoveToFront(c.items[key])
		return true
	}
	// when the key is not in cache, check the capacity, remove back() when the cache is full
	if c.queue.Len() == c.capacity {
		keyToDelete := c.keys[c.queue.Back()]
		delete(c.items, keyToDelete)
		delete(c.keys, c.queue.Back())
		c.queue.Remove(c.queue.Back())
	}
	c.items[key] = c.queue.PushFront(value)
	c.keys[c.items[key]] = key
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// return value if key in items
	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)
		return item.Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
	c.keys = make(map[*ListItem]Key, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		keys:     make(map[*ListItem]Key, capacity),
	}
}
