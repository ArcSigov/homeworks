package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mutex    sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type storage struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	result := false
	c.mutex.Lock()
	item, found := c.items[key]
	if found {
		result = true
		item.Value = storage{key, value}
		c.queue.MoveToFront(item)
	} else {
		if c.queue.Len() == c.capacity {
			value := c.queue.Back().Value.(storage)
			delete(c.items, value.key)
			c.queue.Remove(c.queue.Back())
		}
		c.queue.PushFront(storage{key, value})
		c.items[key] = c.queue.Front()
	}
	c.mutex.Unlock()
	return result
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mutex.Lock()
	item, found := c.items[key]
	c.mutex.Unlock()
	if found {
		c.mutex.Lock()
		c.queue.MoveToFront(item)
		c.mutex.Unlock()
		return item.Value.(storage).value, found
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.mutex.Lock()
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
	c.mutex.Unlock()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
