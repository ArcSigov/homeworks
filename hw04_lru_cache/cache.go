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

// Алгоритм работы кэша:
// - при добавлении элемента:
//     - если элемент присутствует в словаре, то обновить его значение и переместить элемент в начало очереди;
//     - если элемента нет в словаре, то добавить в словарь и в начало очереди
//       (при этом, если размер очереди больше ёмкости кэша,
//       то необходимо удалить последний элемент из очереди и его значение из словаря);
//     - возвращаемое значение - флаг, присутствовал ли элемент в кэше.
// - при получении элемента:
//     - если элемент присутствует в словаре, то переместить элемент в начало очереди и вернуть его значение и true;
//     - если элемента нет в словаре, то вернуть nil и false
//     (работа с кешом похожа на работу с `map`)

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
