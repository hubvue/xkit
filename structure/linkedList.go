package structure

import (
	"fmt"
)

// LinkedLister interface constraints
type LinkedLister[T comparable] interface {
	InsertHead(val T)
	InsertTail(val T)
	InsertBefore(target, val T) bool
	InsertAfter(target, val T) bool
	DeleteHead() (val T, ok bool)
	DeleteTail() (val T, ok bool)
	Delete(val T) bool
	Update(oldV, newV T) bool
	Has(val T) bool
	Head() (val T, ok bool)
	Tail() (val T, ok bool)
	Size() int
	Empty() bool
	Reverse()
}

// LinkedList data structure.
type LinkedList[T comparable] struct {
	head *LinkedNode[T]
	size int
}

// NewLinkedList is the method to create a LinkedList instance.
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// LinkedNode data structure.
type LinkedNode[T comparable] struct {
	Val  T
	Next *LinkedNode[T]
}

// NewLinkedNode is the method to create a LinkedNode instance.
func NewLinkedNode[T comparable](val T, next *LinkedNode[T]) *LinkedNode[T] {
	return &LinkedNode[T]{
		Val:  val,
		Next: next,
	}
}

// The InsertTail method inserts a new element at the tail of the linked list
func (ll *LinkedList[T]) InsertTail(val T) {
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

// The InsertHead method inserts a new element at the head of the linked list
func (ll *LinkedList[T]) InsertHead(val T) {
	newNode := NewLinkedNode(val, ll.head)
	ll.head = newNode
	ll.size++
}

// The InsertAfter method inserts a new element after a specified element
func (ll *LinkedList[T]) InsertAfter(target, val T) bool {
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

// The InsertBefore method inserts a new element before a specified element
func (ll *LinkedList[T]) InsertBefore(target, val T) bool {
	if ll.Empty() {
		return false
	}
	current := ll.head
	if current.Val == target {
		newNode := NewLinkedNode(val, current)
		ll.head = newNode
		ll.size++
		return true
	}
	for current != nil && current.Next != nil {
		if current.Next.Val == target {
			newNode := NewLinkedNode(val, current.Next)
			current.Next = newNode
			ll.size++
			return true
		}
		current = current.Next
	}
	return false
}

// The Update method updates the value of a specified node to a new value
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

// The DeleteHead method deletes the first node of the linked list and returns the node's value
func (ll *LinkedList[T]) DeleteHead() (val T, ok bool) {
	if ll.Empty() {
		return val, false
	}
	val = ll.head.Val
	ll.head = ll.head.Next
	ll.size--
	return val, true
}

// The DeleteTail method deletes the last node of the linked list and returns the node's value
func (ll *LinkedList[T]) DeleteTail() (val T, ok bool) {
	current := ll.head
	if current == nil {
		return val, false
	}
	if current.Next == nil {
		val = current.Val
		ll.head = current.Next
		ll.size--
		return val, true
	}
	for current.Next.Next != nil {
		current = current.Next
	}
	val = current.Next.Val
	current.Next = current.Next.Next
	ll.size--
	return val, true
}

// The Delete method is used to remove the node with the specified value from the linked list
func (ll *LinkedList[T]) Delete(val T) bool {
	current := ll.head
	if current != nil && current.Val == val {
		ll.head = current.Next
		ll.size--
		return true
	}
	for current != nil && current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
			ll.size--
			return true
		}
		current = current.Next
	}
	return false
}

// The Has method is used to check if a node with a specific value exists in the linked list
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

// The Head method is used to obtain the value of the head node of the linked list
func (ll *LinkedList[T]) Head() (val T, ok bool) {
	if ll.head == nil {
		return val, false
	}
	return ll.head.Val, true
}

// The Tail method is used to obtain the value of the tail node of the linked list
func (ll *LinkedList[T]) Tail() (val T, ok bool) {
	if ll.Empty() {
		return val, false
	}
	current := ll.head
	for current.Next != nil {
		current = current.Next
	}
	return current.Val, true
}

// The Size method is used to retrieve the number of nodes in the linked list
func (ll *LinkedList[T]) Size() int {
	return ll.size
}

// The Empty method is used to check whether the linked list is empty
func (ll *LinkedList[T]) Empty() bool {
	return ll.size == 0
}

// The Slice method is used to convert the linked list into a slice
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

// The String method is used to convert the linked list into a string
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

// The Reverse method is used to reverse the linked list
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
