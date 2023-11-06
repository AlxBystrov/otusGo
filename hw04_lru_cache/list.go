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
	len       int
	frontItem *ListItem
	backItem  *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.frontItem
}

func (l *list) Back() *ListItem {
	return l.backItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	// create new item
	item := ListItem{Value: v, Next: l.Front(), Prev: nil}
	// when the list is not empty add the pointer for old front on a new item
	if l.Len() > 0 {
		l.Front().Prev = &item
	} else {
		l.backItem = &item
	}

	l.len++
	l.frontItem = &item
	return &item
}

func (l *list) PushBack(v interface{}) *ListItem {
	// create a new item
	item := ListItem{Value: v, Next: nil, Prev: l.Back()}
	// when the list is not empty add the pointer to old back item
	if l.Len() > 0 {
		l.Back().Next = &item
	} else {
		l.frontItem = &item
	}
	l.backItem = &item
	l.len++
	return &item
}

func (l *list) Remove(i *ListItem) {
	// when the item is front
	if i.Prev == nil && i.Next != nil {
		i.Next.Prev = nil
		l.frontItem = i.Next
	}
	// when the item is not front and not back
	if i.Prev != nil && i.Next != nil {
		i.Prev.Next, i.Next.Prev = i.Next, i.Prev
	}
	// when the item is back
	if i.Prev != nil && i.Next == nil {
		i.Prev.Next = nil
		l.backItem = i.Prev
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	// when the item is not front and not back
	if i.Prev != nil && i.Next != nil {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
		i.Prev = nil
		i.Next = l.frontItem
	}
	// when the item is back
	if i.Prev != nil && i.Next == nil {
		i.Prev.Next = nil
		l.backItem = i.Prev
		i.Prev = nil
		i.Next = l.frontItem
	}
	l.frontItem.Prev = i
	l.frontItem = i
}

func NewList() List {
	return new(list)
}
