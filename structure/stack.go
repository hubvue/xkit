package structure

import (
	"fmt"
)

type stacker[T comparable] interface {
	Push(val T)
	Pop() (val T, ok bool)
	Top() (val T, ok bool)
	Empty() bool
	Size() int
}

type Stack[T comparable] struct {
	head *LinkedNode[T]
	size int
}

func NewStack[T comparable]() Stack[T] {
	return Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	newNode := NewLinkedNode(item, nil)
	newNode.Next = s.head
	s.head = newNode
	s.size++
}

func (s *Stack[T]) Pop() (err error, val T) {
	if s.Empty() {
		err = fmt.Errorf("the stack is empty")
		return err, val
	}
	val = s.head.Val
	s.head = s.head.Next
	s.size--
	return err, val
}

func (s *Stack[T]) Top() (err error, val T) {
	if s.Empty() {
		err = fmt.Errorf("the stack is empty")
		return err, val
	}
	val = s.head.Val
	return err, val
}

func (s *Stack[T]) Empty() bool {
	return s.size == 0
}

func (s *Stack[T]) Size() int {
	return s.size
}
