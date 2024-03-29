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
	len   int
	first *ListItem
	last  *ListItem
}

func NewList() List              { return new(list) }
func (l *list) Front() *ListItem { return l.first }
func (l *list) Back() *ListItem  { return l.last }
func (l *list) Len() int         { return l.len }

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
	l.len++
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
	l.len++
	return l.last
}

func (l *list) Remove(i *ListItem) {
	switch {
	case i == nil || l.len == 0:
		return
	case i == l.first && l.first == l.last:
		l.first = nil
		l.last = l.first
	case i == l.first:
		l.first = i.Next
		l.first.Prev = nil
	case i == l.last:
		l.last = i.Prev
		l.last.Next = nil
	default:
		i.Prev.Next = i.Next
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	switch {
	case i == nil || i == l.first:
		return
	case i == l.last:
		buf := l.last.Prev
		l.last.Next = l.first
		l.first.Prev = l.last
		l.first = l.last
		l.last = buf
		l.first.Prev = nil
		l.last.Next = nil
	default:
		buf := l.first
		l.first.Prev = i
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
		l.first = i
		l.first.Next = buf
		l.first.Prev = nil
	}
}

func (l *ListItem) String() string {
	if v, fl := l.Value.(int); fl {
		return strconv.Itoa(v)
	}
	return ""
}

func (l *list) String() string {
	var out strings.Builder
	out.WriteString("List, len=" + strconv.Itoa(l.Len()) + " [")
	item := l.first
	for i := 0; i < l.len; i++ {
		if i != 0 {
			out.WriteString(",")
		}
		out.WriteString(fmt.Sprint(item.Value))
		item = item.Next
	}
	out.WriteString("]")
	return out.String()
}
