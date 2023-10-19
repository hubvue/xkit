package structure

import "fmt"

type DoublyNode[T comparable] struct {
	Val  T
	Prev *DoublyNode[T]
	Next *DoublyNode[T]
}

func NewDoublyNode[T comparable](val T, prev *DoublyNode[T], next *DoublyNode[T]) *DoublyNode[T] {
	return &DoublyNode[T]{
		Val:  val,
		Prev: prev,
		Next: next,
	}
}

type DoublyLinkedList[T comparable] struct {
	Head *DoublyNode[T]
	Tail *DoublyNode[T]
}

func NewDoublyLinkedList[T comparable]() DoublyLinkedList[T] {
	return DoublyLinkedList[T]{}
}

func (dll *DoublyLinkedList[T]) Add(val T) {
	newNode := NewDoublyNode(val, nil, nil)
	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
		return
	}
	newNode.Prev = dll.Tail
	dll.Tail.Next = newNode
	dll.Tail = newNode
}

func (dll *DoublyLinkedList[T]) Print() {
	current := dll.Head
	for current != nil {
		fmt.Printf("%v -> ", current.Val)
		current = current.Next
	}

	fmt.Println("nil")
}

func (dll *DoublyLinkedList[T]) Reverse() {
	var pre *DoublyNode[T]
	current := dll.Head
	for current != nil {
		next := current.Next
		current.Next = pre
		current.Prev = next
		pre = current
		current = next
	}
	dll.Tail = dll.Head
	dll.Head = pre
}
