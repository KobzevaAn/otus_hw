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
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (lru *lruCache) Set(key Key, value interface{}) bool {
	structItem, inCache := lru.items[key]
	if inCache {
		cache := structItem.Value.(cacheItem)
		cache.value = value
		structItem.Value = cache
		lru.queue.MoveToFront(structItem)
	} else {
		if lru.queue.Len() == lru.capacity {
			delete(lru.items, lru.queue.Back().Value.(cacheItem).key)
			lru.queue.Remove(lru.queue.Back())
		}
		cache := cacheItem{key, value}
		lru.items[key] = lru.queue.PushFront(cache)
	}

	return inCache
}

func (lru *lruCache) Get(key Key) (interface{}, bool) {
	item, inCache := lru.items[key]
	if inCache {
		lru.queue.MoveToFront(item)
		return item.Value.(cacheItem).value, true
	}
	return nil, false
}

func (lru *lruCache) Clear() {
	lru.items = make(map[Key]*ListItem, lru.capacity)
	lru.queue = NewList()
}
