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
	len   int
	front *ListItem
	back  *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{
		Value: v,
		Next:  l.front,
	}

	if l.len == 0 {
		l.back = item
	} else {
		l.front.Prev = item
	}

	l.front = item
	l.len++

	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{
		Value: v,
		Prev:  l.back,
	}

	if l.len == 0 {
		l.front = item
	} else {
		l.back.Next = item
	}

	l.back = item
	l.len++

	return item
}

func (l *list) Remove(i *ListItem) {
	var next = i.Next
	var prev = i.Prev

	if prev == nil {
		l.front = next
	} else {
		prev.Next = next
	}

	if next == nil {
		l.back = prev
	} else {
		next.Prev = prev
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.Len() == 1 || l.Front() == i {
		return
	}

	i.Prev.Next = i.Next
	if i.Next == nil {
		l.back = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	i.Prev = nil
	i.Next = l.front
	l.front = i
}

func NewList() List {
	return new(list)
}
