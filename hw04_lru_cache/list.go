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
	front *ListItem
	back  *ListItem
	len   int
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.front
}

func (l list) Back() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.back
}

func (l *list) init(ext *ListItem) *list {
	l.front = ext
	l.back = ext
	l.len++
	return l
}

func (l *list) PushFront(v interface{}) *ListItem {
	ext := &ListItem{Value: v}
	if l.len == 0 {
		l.init(ext)
	} else {
		ext.Next = l.front
		l.front.Prev = ext
		l.front = ext
		l.len++
	}
	return ext
}

func (l *list) PushBack(v interface{}) *ListItem {
	ext := &ListItem{Value: v}
	if l.len == 0 {
		l.init(ext)
	} else {
		ext.Prev = l.back
		l.back.Next = ext
		l.back = ext
		l.len++
	}
	return ext
}

func NewList() List {
	return new(list)
}

func (l *list) Remove(i *ListItem) {
	switch {
	case i.Prev == nil:
		i.Next.Prev = nil
		l.front = i.Next
	case i.Next == nil:
		i.Prev.Next = nil
		l.back = i.Prev
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	switch i {
	case l.Front():
		l.front = i.Next
		i.Next.Prev = i.Prev
	case l.Back():
		i.Prev.Next = i.Next
		l.back = i.Prev
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	l.front.Prev = i
	i.Next = l.front
	i.Prev = nil
	l.front = i
}
