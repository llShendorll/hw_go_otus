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
	if elem, isOk := l.items[key]; isOk {
		l.queue.MoveToFront(elem)
		elem.Value.(*itemCache).value = value
		return isOk
	}

	item := &itemCache{key: key, value: value}
	l.items[key] = l.queue.PushFront(item)

	if l.capacity < l.queue.Len() {
		elemBack := l.queue.Back()
		l.queue.Remove(elemBack)
		delete(l.items, elemBack.Value.(*itemCache).key)
	}

	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if elem, isOk := l.items[key]; isOk {
		l.queue.MoveToFront(elem)
		return elem.Value.(*itemCache).value, isOk
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
