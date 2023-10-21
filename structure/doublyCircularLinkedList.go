package structure

// DoublyCircularLinkedList data structure
type DoublyCircularLinkedList[T comparable] struct {
	head *DoublyLinkedNode[T]
	tail *DoublyLinkedNode[T]
}

// NewDoublyCircularLinkedList is the method to create a DoublyCircularLinkedList instance
func NewDoublyCircularLinkedList[T comparable]() *DoublyCircularLinkedList[T] {
	return &DoublyCircularLinkedList[T]{}
}
