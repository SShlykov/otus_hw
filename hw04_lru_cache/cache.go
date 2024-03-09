package hw04lrucache

import (
	ll "github.com/SShlykov/otus_hw/hw04_lru_cache/linkedlist"
	"sync"
)

// Key является типом ключа для кэша.
type Key string

// Cache является интерфейсом для работы с кэшем.
type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

// cacheItem является элементом кэша. Он используется для хранения пары ключ-значение.
type cacheItem struct {
	key   Key
	value interface{}
}

// lruCache является реализацией интерфейса Cache. Он представляет собой кэш с алгоритмом LRU.
type lruCache struct {
	capacity int
	queue    ll.List[interface{}]
	items    map[Key]*ll.Node[interface{}]
	mu       sync.Mutex
}

// NewCache создает новый экземпляр кэша с заданной вместимостью.
func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    ll.NewList[interface{}](),
		items:    make(map[Key]*ll.Node[interface{}], capacity),
	}
}

// Get возвращает значение из кэша по ключу.
func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if node, ok := c.items[key]; ok {
		c.queue.MoveToFront(node)
		item := node.Value.(cacheItem)
		return item.value, true
	}
	return nil, false
}

// Set устанавливает значение в кэш по ключу.
func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if node, ok := c.items[key]; ok {
		c.queue.MoveToFront(node)
		node.Value = cacheItem{key, value}
		return true
	}

	if c.queue.Len() == c.capacity {
		last := c.queue.Back()
		lastItem := last.Value.(cacheItem)
		delete(c.items, lastItem.key)
		c.queue.Remove(last)
	}

	node := c.queue.PushFront(cacheItem{key, value})
	c.items[key] = node
	return false
}

// Clear очищает кэш.
func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.queue = ll.NewList[interface{}]()
	c.items = make(map[Key]*ll.Node[interface{}], c.capacity)
}
