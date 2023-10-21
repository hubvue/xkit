package structure

// DoublyLinkedLister interface constraints
type DoublyLinkedLister[T comparable] interface {
	LinkedList[T]
}

// DoublyLinkedNode data structure
type DoublyLinkedNode[T comparable] struct {
	Val  T
	Prev *DoublyLinkedNode[T]
	Next *DoublyLinkedNode[T]
}

// NewDoublyLinkedNode is the method to create a DoublyLinkedNode instance
func NewDoublyLinkedNode[T comparable](val T, prev *DoublyLinkedNode[T], next *DoublyLinkedNode[T]) *DoublyLinkedNode[T] {
	return &DoublyLinkedNode[T]{
		Val:  val,
		Prev: prev,
		Next: next,
	}
}

// DoublyLinkedList data structure
type DoublyLinkedList[T comparable] struct {
	head *DoublyLinkedNode[T]
	tail *DoublyLinkedNode[T]
	size int
}

// NewDoublyLinkedList is the method to create a DoublyLinkedList instance
func NewDoublyLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

// The InsertTail method inserts a new element at the tail of the linked list
func (dll *DoublyLinkedList[T]) InsertTail(val T) {
	newNode := NewDoublyLinkedNode(val, nil, nil)
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
		return
	}
	newNode.Prev = dll.tail
	dll.tail.Next = newNode
	dll.tail = newNode
}

// The InsertHead method inserts a new element at the head of the linked list
func (dll *DoublyLinkedList[T]) InsertHead(val T) {
	newNode := NewDoublyLinkedNode(val, nil, dll.head)
	dll.head.Prev = newNode
	dll.head = newNode
}

// The InsertBefore method inserts a new element before a specified element
func (dll *DoublyLinkedList[T]) InsertBefore(target, val T) bool {

	return false
}

// The InsertAfter method inserts a new element after a specified element
func (dll *DoublyLinkedList[T]) InsertAfter(target, val T) bool {

	return false
}

// The Update method updates the value of a specified node to a new value
func (dll *DoublyLinkedList[T]) Update(oldV, newV T) bool {
	//	TODO:
	return false
}

// The Delete method is used to remove the node with the specified value from the linked list
func (dll *DoublyLinkedList[T]) Delete(val T) bool {
	//	TODO:
	return false
}

// The DeleteHead method deletes the first node of the linked list and returns the node's value
func (dll *DoublyLinkedList[T]) DeleteHead() (val T, ok bool) {
	//	TODO:
	return val, ok
}

// The DeleteTail method deletes the last node of the linked list and returns the node's value
func (dll *DoublyLinkedList[T]) DeleteTail() (val T, ok bool) {
	//	TODO:
	return val, ok
}

// The Has method is used to check if a node with a specific value exists in the linked list
func (dll *DoublyLinkedList[T]) Has(val T) bool {
	//	TODO:
	return false
}

// The Head method is used to obtain the value of the head node of the linked list
func (dll *DoublyLinkedList[T]) Head() (val T, ok bool) {
	//	TODO:
	return val, ok
}

// The Tail method is used to obtain the value of the tail node of the linked list
func (dll *DoublyLinkedList[T]) Tail() (val T, ok bool) {
	//	TODO:
	return val, ok
}

// The Size method is used to retrieve the number of nodes in the linked list
func (dll *DoublyLinkedList[T]) Size() int {
	//	TODO:
	return dll.size
}

// The Empty method is used to check whether the linked list is empty
func (dll *DoublyLinkedList[T]) Empty() bool {
	//	TODO:
	return false
}

// The Slice method is used to convert the linked list into a slice
func (dll *DoublyLinkedList[T]) Slice() []T {
	//	TODO:
	return make([]T, 0)
}

// The Reverse method is used to reverse the linked list
func (dll *DoublyLinkedList[T]) Reverse() {
	var pre *DoublyLinkedNode[T]
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
