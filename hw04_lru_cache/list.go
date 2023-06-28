package hw04lrucache

import (
	"fmt"
	"strconv"
	"strings"
)

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
	count int
	first *ListItem
	last  *ListItem
}

func (l list) Len() int {
	return l.count
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.first == nil {
		l.first = new(ListItem)
		l.first.Value = v
		l.last = l.first
	} else {
		newFirst := new(ListItem)
		newFirst.Value = v
		l.first.Prev = newFirst
		newFirst.Next = l.first
		l.first = newFirst
	}
	l.count++
	return l.first
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.last == nil {
		l.last = new(ListItem)
		l.last.Value = v
		l.first = l.last
	} else {
		newLast := new(ListItem)
		newLast.Value = v
		l.last.Next = newLast
		newLast.Prev = l.last
		l.last = newLast
	}
	l.count++
	return l.last
}

func (l *list) Remove(i *ListItem) {

	if i == nil || l.count == 0 {
		return
	}

	switch {
	case i == l.first && i == l.last:
		l.first = nil
		l.last = l.first
	case i == l.first:
		l.first = i.Next
	case i == l.last:
		l.last = i.Prev
	default:
		i.Prev.Next = i.Next
	}
	l.count--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == nil || i == l.first {
		return
	}

	prev := i.Prev
	prev.Next = i.Next

	if i == l.last {
		l.last = prev
	}

	i.Prev = nil
	i.Next = l.first
	l.first = i
}

func NewList() List {
	return new(list)
}

func (l *ListItem) String() string {
	if v, fl := l.Value.(int); fl {
		return strconv.Itoa(v)
	}
	return ""
}

func (l list) String() string {
	var out strings.Builder
	out.WriteString("List, len=" + strconv.Itoa(l.Len()) + " [")
	item := l.first
	for i := 0; i < l.count; i++ {
		if i != 0 {
			out.WriteString(",")
		}
		out.WriteString(fmt.Sprint(item))
		item = item.Next
	}
	out.WriteString("]")
	return out.String()
}
