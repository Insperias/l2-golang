package main

import "fmt"

type evictionAlg interface {
	evict(c *cache)
}

type fifo struct {
}

func (l *fifo) evict(c *cache) {
	fmt.Println("Убираем запись из кэша по алгоритму FIFO")
}

type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("Убираем запись из экша по алгоритму LRU")
}

type lfu struct {
}

func (l *lfu) evict(c *cache) {
	fmt.Println("Убираем запись из кэша по алгортму LFU")
}

type cache struct {
	storage     map[string]string
	evictionAlg evictionAlg
	capacity    int
	maxCapacity int
}

func initCache(e evictionAlg) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:     storage,
		evictionAlg: e,
		capacity:    0,
		maxCapacity: 3,
	}
}

func (c *cache) setEvictionAlg(e evictionAlg) {
	c.evictionAlg = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	delete(c.storage, key)
}

func (c *cache) evict() {
	c.evictionAlg.evict(c)
	c.capacity--
}

func main() {
	lfu := &lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")

	cache.add("d", "4")

	lru := &lru{}
	cache.setEvictionAlg(lru)

	cache.add("e", "5")

	fifo := &fifo{}
	cache.setEvictionAlg(fifo)

	cache.add("f", "6")
}
