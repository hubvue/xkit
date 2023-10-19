package structure

import "fmt"

type LinkedLister[T comparable] interface {
	Add(val T)
	Delete(val T) bool
	Update(oldV, newV T) bool
	Find(val T) bool
	Head() (val T, ok bool)
	Rear() (val T, ok bool)
	Size() int
	Empty() bool
	Reverse()
}

// LinkedList is a linked list data structure.
type LinkedList[T comparable] struct {
	Head *LinkedNode[T]
}

// NewLinkedList is the method to create a linked list instance.
func NewLinkedList[T comparable]() LinkedList[T] {
	return LinkedList[T]{}
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

// Add method  is inserts an element at the end of a chain table.
func (ll *LinkedList[T]) Add(val T) {
	newNode := NewLinkedNode(val, nil)
	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Delete element from a linked list.
func (ll *LinkedList[T]) Delete(val T) {
	current := ll.Head
	if current != nil && current.Val == val {
		ll.Head = current.Next
		return
	}
	for current != nil {
		if current.Next != nil && current.Next.Val == val {
			current.Next = current.Next.Next
		}
		current = current.Next
		return
	}
}

// Print linked list.
func (ll *LinkedList[T]) Print() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}

	fmt.Println("nil")
}

// Reverse the linked list
func (ll *LinkedList[T]) Reverse() {
	var pre *LinkedNode[T]
	current := ll.Head

	for current != nil {
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}
	ll.Head = pre
}
