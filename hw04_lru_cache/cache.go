package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	item, finded := c.items[key]
	if finded {
		item.Value = value
		c.queue.MoveToFront(item)
	} else {
		if c.queue.Len() == c.capacity {
			delete(c.items, key)
			c.queue.Remove(c.queue.Back())
		}
		c.queue.PushFront(value)
		c.items[key] = c.queue.Front()
	}
	return finded
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	item, finded := c.items[key]
	if finded {
		c.queue.MoveToFront(item)
		return item.Value, finded
	}
	return nil, false
}

func (c *lruCache) Clear() {

}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
