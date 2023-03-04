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

type itemCache struct {
	key   Key
	value interface{}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	elem := l.items[key]
	if elem != nil {
		l.queue.MoveToFront(elem)
		elem.Value.(*itemCache).value = value
		return true
	}

	item := &itemCache{key: key, value: value}
	l.items[key] = l.queue.PushFront(item)
	if l.capacity < l.queue.Len() {
		l.queue.Remove(l.queue.Back())
	}

	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	elem := l.items[key]
	if elem != nil {
		l.queue.MoveToFront(elem)
		return elem.Value.(*itemCache).value, true
	}

	return nil, false
}

func (l *lruCache) Clear() {
	l.items = make(map[Key]*ListItem, l.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
