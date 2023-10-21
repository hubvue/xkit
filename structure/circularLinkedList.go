package structure

// CircularLinkedList data structure
type CircularLinkedList[T comparable] struct {
	head *LinkedNode[T]
	tail *LinkedNode[T]
}

// NewCircularLinkedList is the method to create a CircularLinkedList instance
func NewCircularLinkedList[T comparable]() *CircularLinkedList[T] {
	return &CircularLinkedList[T]{}
}
