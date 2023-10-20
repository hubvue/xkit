package structure

// DoublyLinkedLister interface constraints
type DoublyLinkedLister[T comparable] interface {
	LinkedList[T]
	Tail() (val T, ok bool)
	AddHead(val T)
	DeleteHead() (val T, ok bool)
	DeleteTail() (val T, ok bool)
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

// Add method is inserts an element at the end of a linked list.
func (dll *DoublyLinkedList[T]) Add(val T) {
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

// AddHead method is inserts an element at the head of a linked list
func (dll *DoublyLinkedList[T]) AddHead(val T) {
	newNode := NewDoublyLinkedNode(val, nil, dll.head)
	dll.head.Prev = newNode
	dll.head = newNode
}

func (dll *DoublyLinkedList[T]) Insert(target, val T) bool {
	//	TODO:
	return false
}

func (dll *DoublyLinkedList[T]) Delete(val T) bool {
	//	TODO:
	return false
}

func (dll *DoublyLinkedList[T]) DeleteHead() (val T, ok bool) {
	//	TODO:
	return val, ok
}

func (dll *DoublyLinkedList[T]) DeleteTail() (val T, ok bool) {
	//	TODO:
	return val, ok
}

func (dll *DoublyLinkedList[T]) Update(oldV, newV T) bool {
	//	TODO:
	return false
}

func (dll *DoublyLinkedList[T]) Has(val T) bool {
	//	TODO:
	return false
}

func (dll *DoublyLinkedList[T]) Head() (val T, ok bool) {
	//	TODO:
	return val, ok
}

func (dll *DoublyLinkedList[T]) Tail() (val T, ok bool) {
	//	TODO:
	return val, ok
}

func (dll *DoublyLinkedList[T]) Size() int {
	//	TODO:
	return dll.size
}

func (dll *DoublyLinkedList[T]) Empty() bool {
	//	TODO:
	return false
}

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
