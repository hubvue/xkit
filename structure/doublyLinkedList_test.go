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

}

func TestDoublyLinkedList_InsertTail(t *testing.T) {

}

func TestDoublyLinkedList_InsertBefore(t *testing.T) {

}

func TestDoublyLinkedList_InsertAfter(t *testing.T) {

}

func TestDoublyLinkedList_Update(t *testing.T) {

}

func TestDoublyLinkedList_Delete(t *testing.T) {

}

func TestDoublyLinkedList_DeleteHead(t *testing.T) {

}

func TestDoublyLinkedList_DeleteTail(t *testing.T) {

}

func TestDoublyLinkedList_Has(t *testing.T) {

}

func TestDoublyLinkedList_Head(t *testing.T) {

}

func TestDoublyLinkedList_Tail(t *testing.T) {

}

func TestDoublyLinkedList_Size(t *testing.T) {

}

func TestDoublyLinkedList_Empty(t *testing.T) {

}

func TestDoublyLinkedList_Slice(t *testing.T) {

}

func TestDoublyLinkedList_Reverse(t *testing.T) {

}
