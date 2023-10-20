package structure

import "fmt"

type DoublyLinkedLister[T comparable] interface {
	LinkedList[T]
}

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
	head *DoublyNode[T]
	tail *DoublyNode[T]
}

func NewDoublyLinkedList[T comparable]() DoublyLinkedList[T] {
	return DoublyLinkedList[T]{}
}

func (dll *DoublyLinkedList[T]) Add(val T) {
	newNode := NewDoublyNode(val, nil, nil)
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
		return
	}
	newNode.Prev = dll.tail
	dll.tail.Next = newNode
	dll.tail = newNode
}

func (dll *DoublyLinkedList[T]) Print() {
	current := dll.head
	for current != nil {
		fmt.Printf("%v -> ", current.Val)
		current = current.Next
	}

	fmt.Println("nil")
}

func (dll *DoublyLinkedList[T]) Reverse() {
	var pre *DoublyNode[T]
	current := dll.head
	for current != nil {
		next := current.Next
		current.Next = pre
		current.Prev = next
		pre = current
		current = next
	}
	dll.tail = dll.head
	dll.head = pre
}
