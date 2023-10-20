package structure

import (
	"fmt"
)

type LinkedLister[T comparable] interface {
	Add(val T)
	Insert(target, val T) bool
	Delete(val T) bool
	Update(oldV, newV T) bool
	Has(val T) bool
	Head() (val T, ok bool)
	Size() int
	Empty() bool
	Reverse()
}

// LinkedList is a linked list data structure.
type LinkedList[T comparable] struct {
	head *LinkedNode[T]
	size int
}

// NewLinkedList is the method to create a linked list instance.
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// LinkedNode is a linked list node data structure.
type LinkedNode[T comparable] struct {
	Val  T
	Next *LinkedNode[T]
}

// NewLinkedNode is the method to create a linked list node instance.
func NewLinkedNode[T comparable](val T, next *LinkedNode[T]) *LinkedNode[T] {
	return &LinkedNode[T]{
		Val:  val,
		Next: next,
	}
}

// Add method is inserts an element at the end of a chain table.
func (ll *LinkedList[T]) Add(val T) {
	newNode := NewLinkedNode(val, nil)
	ll.size++
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// The Insert method inserts after an existing node in a linked list.
func (ll *LinkedList[T]) Insert(target, val T) bool {
	current := ll.head
	for current != nil {
		if current.Val == target {
			next := current.Next
			newNode := NewLinkedNode[T](val, next)
			current.Next = newNode
			ll.size++
			return true
		}
		current = current.Next
	}
	return false
}

// The Update method is used to update the value of an existing node in a linked list.
func (ll *LinkedList[T]) Update(oldV, newV T) bool {
	current := ll.head
	for current != nil {
		if current.Val == oldV {
			current.Val = newV
			return true
		}
		current = current.Next
	}
	return false
}

// The Has method is used to find out if a value exists in a linked list.
func (ll *LinkedList[T]) Has(val T) bool {
	current := ll.head
	for current != nil {
		if current.Val == val {
			return true
		}
		current = current.Next
	}
	return false
}

// The Head method is used to get the value of the head node in a linked list.
func (ll *LinkedList[T]) Head() (val T, ok bool) {
	if ll.head == nil {
		return val, false
	}
	return ll.head.Val, true
}

// The Size method is used to get the number of nodes in a linked list.
func (ll *LinkedList[T]) Size() int {
	return ll.size
}

// The Empty method is used to determine if a linked list is empty.
func (ll *LinkedList[T]) Empty() bool {
	return ll.size == 0
}

// Delete element from a linked list.
func (ll *LinkedList[T]) Delete(val T) bool {
	current := ll.head
	if current != nil && current.Val == val {
		ll.head = current.Next
		ll.size--
		return true
	}
	for current != nil {
		if current.Next != nil && current.Next.Val == val {
			current.Next = current.Next.Next
			ll.size--
			return true
		}
		current = current.Next
	}
	return false
}

func (ll *LinkedList[T]) Slice() []T {
	slice := make([]T, ll.size)
	current := ll.head
	i := 0
	for current != nil {
		slice[i] = current.Val
		current = current.Next
	}
	return slice
}

// The String method is used to convert a linked list into a string output
func (ll *LinkedList[T]) String() string {
	current := ll.head
	var s string
	for current != nil {
		s = s + fmt.Sprintf("%v -> ", current.Val)
		current = current.Next
	}
	s = s + "nil"
	return s
}

// Reverse the linked list
func (ll *LinkedList[T]) Reverse() {
	var pre *LinkedNode[T]
	current := ll.head

	for current != nil {
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}
	ll.head = pre
}
