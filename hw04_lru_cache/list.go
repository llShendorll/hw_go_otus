package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	size  int
	front *ListItem
	back  *ListItem
}

func (l *list) Len() int {
	return l.size
}

func (l *list) Front() *ListItem {
	if l.size == 0 {
		return nil
	}

	return l.front
}

func (l *list) Back() *ListItem {
	if l.size == 0 {
		return nil
	}

	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := ListItem{Value: v, Next: l.front}

	if item.Next != nil {
		l.front.Prev = &item
	}
	if l.back == nil {
		l.back = &item
	}

	l.front = &item

	l.size++
	return &item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := ListItem{Value: v, Prev: l.back}

	if item.Prev != nil {
		l.back.Next = &item
	} else {
		l.front = &item
	}

	l.back = &item

	l.size++
	return &item
}

func (l *list) Remove(i *ListItem) {
	if l.size == 0 {
		return
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.front = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.back = i.Next
	}

	i = nil

	l.size--
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}

	if i.Next == nil {
		i.Prev.Next = nil
		l.back = i.Prev
	} else {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	i.Next = l.front
	i.Next.Prev = i
	i.Prev = nil
	l.front = i
}

func NewList() List {
	return new(list)
}
