package structure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList[int]()
	assert.Equal(t, ll.size, 0)
	assert.Nil(t, ll.head)
	ll.InsertTail(1)
	assert.Equal(t, ll.size, 1)
	assert.NotNil(t, ll.head)
}

func TestNewLinkedNode(t *testing.T) {
	n2 := NewLinkedNode[int](2, nil)
	n1 := NewLinkedNode[int](1, n2)
	assert.Equal(t, n1.Val, 1)
	assert.Equal(t, n2.Val, 2)
	assert.Equal(t, n1.Next, n2)
}

func TestLinkedList_InsertTail(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.InsertTail(1)
	ll.InsertTail(2)
	assert.Equal(t, ll.size, 2)
	assert.NotNil(t, ll.head)
	assert.Equal(t, ll.head.Val, 1)
	assert.NotNil(t, ll.head.Next)
	assert.Equal(t, ll.head.Next.Val, 2)
}

func TestLinkedList_InsertHead(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.InsertHead(1)
	assert.NotNil(t, ll.head)
	assert.Equal(t, ll.head.Val, 1)
	ll.InsertHead(2)
	assert.Equal(t, ll.head.Val, 2)
}

func TestLinkedList_InsertBefore(t *testing.T) {
	ll := NewLinkedList[int]()
	ok := ll.InsertBefore(1, 2)
	assert.Equal(t, ok, false)
	ll.InsertTail(1)
	assert.NotNil(t, ll.head)
	assert.Equal(t, ll.head.Val, 1)
	ok = ll.InsertBefore(1, 2)
	assert.Equal(t, ok, true)
	assert.Equal(t, ll.head.Val, 2)
	assert.NotNil(t, ll.head.Next)
	assert.Equal(t, ll.head.Next.Val, 1)
	ok = ll.InsertBefore(1, 3)
	assert.Equal(t, ok, true)
	val, _ := ll.Tail()
	assert.Equal(t, val, 1)
	ok = ll.InsertBefore(4, 5)
	assert.Equal(t, ok, false)
}

func TestLinkedList_InsertAfter(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.InsertTail(1)
	ok := ll.InsertAfter(1, 2)
	assert.Equal(t, ll.size, 2)
	assert.Equal(t, ok, true)
	ok = ll.InsertAfter(3, 4)
	assert.Equal(t, ok, false)
	assert.Equal(t, ll.size, 2)
}

func TestLinkedList_Update(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.InsertTail(1)
	ll.InsertTail(2)
	ok := ll.Update(1, 3)
	assert.Equal(t, ok, true)
	ok = ll.Update(4, 5)
	assert.Equal(t, ok, false)
	val, _ := ll.Head()
	assert.Equal(t, val, 3)
}

func TestLinkedList_DeleteHead(t *testing.T) {
	ll := NewLinkedList[int]()
	_, ok := ll.DeleteHead()
	assert.Equal(t, ok, false)
	ll.InsertHead(1)
	val, ok := ll.DeleteHead()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 1)
	_, ok = ll.DeleteHead()
	assert.Equal(t, ok, false)
}

func TestLinkedList_DeleteTail(t *testing.T) {
	ll := NewLinkedList[int]()
	_, ok := ll.DeleteTail()
	assert.Equal(t, ok, false)
	ll.InsertHead(1)
	ll.InsertHead(2)
	val, ok := ll.DeleteTail()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 1)
	ll.InsertHead(3)
	ll.InsertHead(4)
	ll.InsertHead(5)
	val, ok = ll.DeleteTail()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 2)
	_, ok = ll.DeleteTail()
	_, ok = ll.DeleteTail()
	_, ok = ll.DeleteTail()
	_, ok = ll.DeleteTail()
	assert.Equal(t, ok, false)
}

func TestLinkedList_Delete(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.InsertTail(1)
	ll.InsertTail(2)
	ok := ll.Delete(1)
	assert.Equal(t, ok, true)
	assert.Equal(t, ll.size, 1)
	ll.InsertTail(3)
	ll.InsertTail(4)
	ll.InsertTail(5)
	ll.InsertTail(6)
	ll.InsertTail(7)
	ok = ll.Delete(8)
	assert.Equal(t, ok, false)
}

func TestLinkedList_Empty(t *testing.T) {
	ll := NewLinkedList[int]()
	ok := ll.Empty()
	assert.Equal(t, ok, true)
	ll.InsertTail(1)
	ok = ll.Empty()
	assert.Equal(t, ok, false)
}

func TestLinkedList_Has(t *testing.T) {
	ll := NewLinkedList[int]()
	ok := ll.Has(1)
	assert.Equal(t, ok, false)
	ll.InsertTail(1)
	ll.InsertTail(2)
	ok = ll.Has(2)
	assert.Equal(t, ok, true)
}

func TestLinkedList_Head(t *testing.T) {
	ll := NewLinkedList[int]()
	_, ok := ll.Head()
	assert.Equal(t, ok, false)
	ll.InsertTail(1)
	val, ok := ll.Head()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 1)
	ll.InsertTail(2)
	val, _ = ll.Head()
	assert.Equal(t, val, 1)
}

func TestLinkedList_Tail(t *testing.T) {
	ll := NewLinkedList[int]()
	_, ok := ll.Tail()
	assert.Equal(t, ok, false)
	ll.InsertTail(1)
	val, ok := ll.Tail()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 1)
	ll.InsertTail(2)
	val, ok = ll.Tail()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 2)
}

func TestLinkedList_Reverse(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.InsertTail(1)
	ll.InsertTail(2)
	ll.InsertTail(3)
	ll.Reverse()
	val, _ := ll.Head()
	assert.Equal(t, val, 3)
}

func TestLinkedList_Size(t *testing.T) {
	ll := NewLinkedList[int]()
	size := ll.Size()
	assert.Equal(t, size, 0)
	ll.InsertTail(1)
	size = ll.Size()
	assert.Equal(t, size, 1)
	ll.InsertAfter(1, 2)
	size = ll.Size()
	assert.Equal(t, size, 2)
	ll.Delete(2)
	size = ll.Size()
	assert.Equal(t, size, 1)
	ll.InsertHead(1)
	assert.Equal(t, ll.Size(), 2)
	ll.DeleteHead()
	assert.Equal(t, ll.Size(), 1)
	ll.DeleteTail()
	assert.Equal(t, ll.Size(), 0)
}

func TestLinkedList_String(t *testing.T) {
	ll := NewLinkedList[int]()
	s := ll.String()
	assert.Equal(t, s, "nil")
	ll.InsertTail(1)
	ll.InsertTail(2)
	s = ll.String()
	expected := "1 -> 2 -> nil"
	assert.Equal(t, s, expected)
}

func TestLinkedList_Slice(t *testing.T) {
	ll := NewLinkedList[int]()
	s := ll.Slice()
	assert.Equal(t, len(s), 0)
	ll.InsertTail(0)
	ll.InsertTail(1)
	ll.InsertTail(2)
	s = ll.Slice()
	assert.Equal(t, len(s), ll.Size())
}

func testTypeHandler[T comparable](t *testing.T, expected T) {
	ll := NewLinkedList[T]()
	ll.InsertHead(expected)
	val, _ := ll.Head()
	assert.Equal(t, val, expected)
}
func TestLinkedList_Type(t *testing.T) {
	testTypeHandler[float32](t, 1.2)
	testTypeHandler[float64](t, 1.23)
	testTypeHandler[int](t, 1)
	testTypeHandler[int16](t, 1111)
	testTypeHandler[int32](t, 111111)
	testTypeHandler[int64](t, 11111)
	testTypeHandler[string](t, "test")
}
