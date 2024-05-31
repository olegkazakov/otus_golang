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

type cacheItem struct {
	Key   Key
	Value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if item, ok := c.items[key]; ok {
		item.Value = cacheItem{Key: key, Value: value}
		c.queue.MoveToFront(item)

		return true
	}

	if c.capacity == c.queue.Len() {
		back := c.queue.Back()
		delete(c.items, back.Value.(cacheItem).Key)
		c.queue.Remove(back)
	}

	c.queue.PushFront(cacheItem{Key: key, Value: value})
	c.items[key] = c.queue.Front()

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)

		return item.Value.(cacheItem).Value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = NewList()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
