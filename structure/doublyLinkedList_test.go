package structure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDoublyLinkedList(t *testing.T) {
	dll := NewDoublyLinkedList[int]()
	assert.Equal(t, dll.size, 0)
	assert.Nil(t, dll.head)
	assert.Nil(t, dll.tail)
	dll.InsertHead(1)
	assert.Equal(t, dll.size, 1)
	assert.NotNil(t, dll.head)
	assert.NotNil(t, dll.tail)
}

func TestNewDoublyLinkedNode(t *testing.T) {
	prev := NewDoublyLinkedNode(1, nil, nil)
	assert.Equal(t, prev.Val, 1)
	assert.Nil(t, prev.Prev)
	assert.Nil(t, prev.Next)
	next := NewDoublyLinkedNode(3, nil, nil)
	cur := NewDoublyLinkedNode(2, prev, next)
	prev.Next = cur
	next.Prev = cur
	assert.Equal(t, cur.Prev, prev)
	assert.Equal(t, cur.Next, next)
}

func TestDoublyLinkedList_InsertHead(t *testing.T) {
	dll := NewDoublyLinkedList[int]()
	assert.NotNil(t, dll)
	assert.Nil(t, dll.head)
	assert.Nil(t, dll.tail)
	dll.InsertHead(1)
	assert.NotNil(t, dll.head)
	assert.NotNil(t, dll.tail)
	assert.Equal(t, dll.head.Val, 1)
	assert.Equal(t, dll.tail.Val, 1)
	dll.InsertHead(2)
	assert.Equal(t, dll.head.Val, 2)
	assert.Equal(t, dll.tail.Val, 1)
}

func TestDoublyLinkedList_InsertTail(t *testing.T) {
	dll := NewDoublyLinkedList[int]()
	assert.NotNil(t, dll)
	assert.Nil(t, dll.head)
	assert.Nil(t, dll.tail)
	dll.InsertTail(1)
	assert.NotNil(t, dll.head)
	assert.NotNil(t, dll.tail)
	assert.Equal(t, dll.head.Val, 1)
	assert.Equal(t, dll.tail.Val, 1)
	dll.InsertTail(2)
	assert.Equal(t, dll.head.Val, 1)
	assert.Equal(t, dll.tail.Val, 2)
}

func TestDoublyLinkedList_InsertBefore(t *testing.T) {
	dll := NewDoublyLinkedList[int]()
	ok := dll.InsertBefore(1, 2)
	assert.Equal(t, ok, false)
	dll.InsertHead(1)
	dll.InsertBefore(1, 2)
	assert.Equal(t, dll.head.Val, 2)
	dll.InsertBefore(1, 3)
	assert.Equal(t, dll.head.Next.Val, 3)
	dll.InsertBefore(3, 4)
	assert.Equal(t, dll.head.Next.Val, 4)
	dll.InsertBefore(1, 5)
	assert.Equal(t, dll.tail.Prev.Val, 5)
	ok = dll.InsertBefore(6, 7)
	assert.Equal(t, ok, false)

}

func TestDoublyLinkedList_InsertAfter(t *testing.T) {
	dll := NewDoublyLinkedList[int]()
	ok := dll.InsertAfter(1, 2)
	assert.Equal(t, ok, false)
	dll.InsertHead(1)
	dll.InsertAfter(1, 2)
	assert.Equal(t, dll.tail.Val, 2)
	dll.InsertAfter(1, 3)
	assert.Equal(t, dll.head.Next.Val, 3)
	dll.InsertAfter(3, 4)
	assert.Equal(t, dll.tail.Prev.Val, 4)
	dll.InsertAfter(3, 5)
	dll.InsertAfter(5, 6)
	dll.InsertAfter(4, 7)
	assert.Equal(t, dll.tail.Prev.Val, 7)
	dll.InsertAfter(2, 10)
	assert.Equal(t, dll.tail.Val, 10)
	ok = dll.InsertAfter(8, 9)
	assert.Equal(t, ok, false)
}

func TestDoublyLinkedList_Update(t *testing.T) {
	dll := NewDoublyLinkedList[int]()
	ok := dll.Update(1, 2)
	assert.Equal(t, ok, false)
	dll.InsertHead(1)
	ok = dll.Update(1, 2)
	assert.Equal(t, ok, true)
	dll.InsertTail(3)
	dll.InsertTail(4)
	dll.InsertTail(5)
	dll.InsertTail(6)
	ok = dll.Update(6, 6*2)
	assert.Equal(t, ok, true)
	ok = dll.Update(4, 4*2)
	assert.Equal(t, ok, true)
	ok = dll.Update(7, 8*2)
	assert.Equal(t, ok, false)
}

func TestDoublyLinkedList_Delete(t *testing.T) {
	dll := NewDoublyLinkedList[int]()
	ok := dll.Delete(1)
	assert.Equal(t, ok, false)
	dll.InsertTail(1)
	ok = dll.Delete(1)
	assert.Equal(t, ok, true)
	dll.InsertTail(2)
	dll.InsertTail(3)
	dll.InsertTail(4)
	dll.InsertTail(5)
	dll.InsertTail(6)
	dll.InsertTail(7)
	ok = dll.Delete(2)
	assert.Equal(t, ok, true)
	ok = dll.Delete(7)
	assert.Equal(t, ok, true)
	ok = dll.Delete(4)
	assert.Equal(t, ok, true)
	dll.InsertTail(11)
	dll.InsertTail(12)
	ok = dll.Delete(11)
	assert.Equal(t, ok, true)
	ok = dll.Delete(10)
	assert.Equal(t, ok, false)
}

func TestDoublyLinkedList_DeleteHead(t *testing.T) {
	dll := NewDoublyLinkedList[int]()
	_, ok := dll.DeleteHead()
	assert.Equal(t, ok, false)
	dll.InsertTail(1)
	dll.InsertTail(2)
	val, ok := dll.DeleteHead()
	assert.Equal(t, val, 1)
	assert.Equal(t, ok, true)
	val, ok = dll.DeleteHead()
	assert.Equal(t, val, 2)
	assert.Equal(t, ok, true)
	val, ok = dll.DeleteHead()
	assert.Equal(t, ok, false)
}

func TestDoublyLinkedList_DeleteTail(t *testing.T) {
	dll := NewDoublyLinkedList[int]()
	_, ok := dll.DeleteTail()
	assert.Equal(t, ok, false)
	dll.InsertTail(1)
	dll.InsertTail(2)
	val, ok := dll.DeleteTail()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 2)
	val, ok = dll.DeleteTail()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 1)
	val, ok = dll.DeleteTail()
	assert.Equal(t, ok, false)
}

func TestDoublyLinkedList_Has(t *testing.T) {
	dll := NewDoublyLinkedList[int]()
	ok := dll.Has(1)
	assert.Equal(t, ok, false)
	dll.InsertTail(1)
	dll.InsertTail(2)
	dll.InsertTail(3)
	dll.InsertTail(4)
	dll.InsertTail(5)
	ok = dll.Has(1)
	assert.Equal(t, ok, true)
	ok = dll.Has(5)
	assert.Equal(t, ok, true)
	ok = dll.Has(3)
	assert.Equal(t, ok, true)
	ok = dll.Has(6)
	assert.Equal(t, ok, false)
}

func TestDoublyLinkedList_Head(t *testing.T) {
	dll := NewDoublyLinkedList[int]()
	_, ok := dll.Head()
	assert.Equal(t, ok, false)
	dll.InsertHead(1)
	val, ok := dll.Head()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 1)
}

func TestDoublyLinkedList_Tail(t *testing.T) {
	dll := NewDoublyLinkedList[int]()
	_, ok := dll.Tail()
	assert.Equal(t, ok, false)
	dll.InsertTail(1)
	val, ok := dll.Tail()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 1)
}

func TestDoublyLinkedList_Size(t *testing.T) {
	dll := NewDoublyLinkedList[int]()
	assert.Equal(t, dll.Size(), 0)
	dll.InsertHead(1)
	assert.Equal(t, dll.Size(), 1)
	dll.InsertTail(2)
	assert.Equal(t, dll.Size(), 2)
	dll.InsertAfter(1, 3)
	assert.Equal(t, dll.Size(), 3)
	dll.InsertBefore(1, 4)
	assert.Equal(t, dll.Size(), 4)
	dll.DeleteHead()
	assert.Equal(t, dll.Size(), 3)
	dll.DeleteTail()
	assert.Equal(t, dll.Size(), 2)
	dll.Delete(3)
	assert.Equal(t, dll.Size(), 1)
}

func TestDoublyLinkedList_Empty(t *testing.T) {
	dll := NewDoublyLinkedList[int]()
	assert.Equal(t, dll.Empty(), true)
	dll.InsertHead(1)
	assert.Equal(t, dll.Empty(), false)
	dll.DeleteHead()
	assert.Equal(t, dll.Empty(), true)
}

func TestDoublyLinkedList_Slice(t *testing.T) {
	dll := NewDoublyLinkedList[int]()
	s := dll.Slice()
	assert.Equal(t, len(s), len(dll.Slice()))
	dll.InsertHead(1)
	dll.InsertHead(1)
	dll.InsertHead(1)
	dll.InsertHead(1)
	s = dll.Slice()
	assert.Equal(t, len(s), len(dll.Slice()))
}

func TestDoublyLinkedList_Reverse(t *testing.T) {
	dll := NewDoublyLinkedList[int]()
	dll.Reverse()
	assert.Nil(t, dll.head)
	dll.InsertHead(1)
	dll.InsertHead(2)
	dll.InsertHead(3)
	dll.InsertHead(4)
	dll.Reverse()
	assert.Equal(t, dll.head.Val, 1)
}
