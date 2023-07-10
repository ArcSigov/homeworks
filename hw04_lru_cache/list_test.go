package hw04lrucache

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}

func TestHardList(t *testing.T) {
	t.Run("Equal back and front", func(t *testing.T) {
		l := NewList()
		l.PushBack(1)
		require.Equal(t, 1, l.Front().Value) // 1.
		require.Equal(t, 1, l.Back().Value)  // 1.
	})

	t.Run("Remove and move nil item", func(t *testing.T) {
		l := NewList()
		l.PushBack(1)
		l.Remove(nil)
		l.MoveToFront(nil)
		require.Equal(t, 1, l.Front().Value)
		require.Equal(t, 1, l.Len())
	})

	t.Run("Remove at Front", func(t *testing.T) {
		l := NewList()
		l.PushBack(1)
		l.PushBack(2)
		l.Remove(l.Front())
		require.Equal(t, 2, l.Front().Value)
		require.Equal(t, 1, l.Len())
	})

	t.Run("Remove at Back", func(t *testing.T) {
		l := NewList()
		l.PushBack(1)
		l.PushBack(2)
		l.Remove(l.Back())
		require.Equal(t, 1, l.Front().Value)
		require.Equal(t, 1, l.Len())
	})

	t.Run("Remove once", func(t *testing.T) {
		l := NewList()
		l.PushBack(1)
		l.Remove(l.Back())
		require.Equal(t, l.Back(), l.Front())
		require.Equal(t, 0, l.Len())
	})

	t.Run("Middle move", func(t *testing.T) {
		l := NewList()
		l.PushBack(2)
		middle := l.PushBack(1)
		l.PushBack(3)
		l.MoveToFront(middle)
		require.Equal(t, 1, l.Front().Value)
		require.Equal(t, 3, l.Len())
	})

	t.Run("Move back", func(t *testing.T) {
		l := NewList()
		l.PushBack(2)
		l.PushBack(1)
		l.MoveToFront(l.Back())
		require.Equal(t, 1, l.Front().Value)
		require.Equal(t, 2, l.Len())
	})
}

func TestPrintList(t *testing.T) {
	l := NewList()
	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}
	t.Run("Full List", func(t *testing.T) {
		expected := "List, len=10 [0,1,2,3,4,5,6,7,8,9]"
		require.Equal(t, expected, fmt.Sprint(l))
	})

	t.Run("Item List", func(t *testing.T) {
		expected := "9"
		require.Equal(t, expected, fmt.Sprint(l.Back()))
	})

	t.Run("Item nil", func(t *testing.T) {
		item := new(ListItem)
		item.Value = "a"
		require.Equal(t, "", fmt.Sprint(item))
	})
}
