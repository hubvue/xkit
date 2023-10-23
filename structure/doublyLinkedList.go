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
	dll.size++
}

// The InsertHead method inserts a new element at the head of the linked list
func (dll *DoublyLinkedList[T]) InsertHead(val T) {
	newNode := NewDoublyLinkedNode(val, nil, dll.head)
	dll.head.Prev = newNode
	dll.head = newNode
	dll.size++
}

// The InsertBefore method inserts a new element before a specified element
func (dll *DoublyLinkedList[T]) InsertBefore(target, val T) bool {
	if dll.Empty() {
		return false
	}
	head := dll.head
	if head.Val == target {
		newNode := NewDoublyLinkedNode(val, nil, head)
		head.Prev = newNode
		dll.head = newNode
		dll.size++
		return true
	}
	tail := dll.tail
	for head.Prev != tail && head != tail {
		if head.Next.Val == target {
			newNode := NewDoublyLinkedNode(val, head, head.Next)
			head.Next.Prev = newNode
			head.Next = newNode
			dll.size++
			return true
		}
		if tail.Val == target {
			newNode := NewDoublyLinkedNode(val, tail.Prev, tail)
			tail.Prev.Next = newNode
			tail.Prev = newNode
			dll.size++
			return true
		}
		head = head.Next
		tail = tail.Prev
	}
	return false
}

// The InsertAfter method inserts a new element after a specified element
func (dll *DoublyLinkedList[T]) InsertAfter(target, val T) bool {
	if dll.Empty() {
		return false
	}
	head := dll.head
	if head.Val == target {
		newNode := NewDoublyLinkedNode(val, head, head.Next)
		if head.Next != nil {
			head.Next.Prev = newNode
		} else {
			dll.tail = newNode
		}
		head.Next = newNode
		dll.size++
		return true
	}
	tail := dll.tail
	for head.Prev != tail && head != tail {
		if head.Val == target {
			newNode := NewDoublyLinkedNode(val, head, head.Next)
			head.Next.Prev = newNode
			head.Next = newNode
			dll.size++
			return true
		}
		if tail.Val == target {
			newNode := NewDoublyLinkedNode(val, tail, tail.Next)
			if tail.Next != nil {
				tail.Next.Prev = newNode
			} else {
				dll.tail = newNode
			}
			tail.Next = newNode
			dll.size++
			return true
		}
		head = head.Next
		tail = tail.Prev
	}

	return false
}

// The Update method updates the value of a specified node to a new value
func (dll *DoublyLinkedList[T]) Update(oldV, newV T) bool {
	if dll.Empty() {
		return false
	}
	head := dll.head
	if head.Val == oldV {
		head.Val = newV
		return true
	}
	tail := dll.tail
	for head.Prev != tail && head != tail {
		if head.Val == oldV {
			head.Val = newV
			return true
		}
		if tail.Val == oldV {
			tail.Val = newV
			return true
		}
		head = head.Next
		tail = tail.Prev
	}
	return false
}

// The Delete method is used to remove the node with the specified value from the linked list
func (dll *DoublyLinkedList[T]) Delete(val T) bool {
	if dll.Empty() {
		return false
	}
	head := dll.head
	if head.Val == val {
		dll.head = dll.head.Next
		if dll.head == nil {
			dll.tail = nil
		}
		dll.size--
		return true
	}
	tail := dll.tail
	for head.Prev != tail && head != tail {
		if head.Val == val {
			head.Next.Prev = head.Prev
			head.Prev.Next = head.Next
			head.Next = nil
			head.Prev = nil
			dll.size--
			return true
		}
		if tail.Val == val {
			if tail.Next != nil {
				tail.Next.Prev = tail.Prev
			} else {
				dll.tail = tail.Prev
			}
			tail.Prev.Next = tail.Next
			tail.Next = nil
			tail.Prev = nil
			dll.size--
			return true
		}
	}
	return false
}

// The DeleteHead method deletes the first node of the linked list and returns the node's value
func (dll *DoublyLinkedList[T]) DeleteHead() (val T, ok bool) {
	if dll.Empty() {
		return val, false
	}
	val = dll.head.Val
	next := dll.head.Next
	dll.head.Next = nil
	dll.head = next
	if dll.head != nil {
		dll.head.Prev = nil
	} else {
		dll.tail = nil
	}
	dll.size--
	return val, true
}

// The DeleteTail method deletes the last node of the linked list and returns the node's value
func (dll *DoublyLinkedList[T]) DeleteTail() (val T, ok bool) {
	if dll.Empty() {
		return val, false
	}
	val = dll.tail.Val
	prev := dll.tail.Prev
	dll.tail.Prev = nil
	dll.tail = prev
	if dll.tail != nil {
		dll.tail.Next = nil
	} else {
		dll.head = nil
	}
	dll.size--
	return val, true
}

// The Has method is used to check if a node with a specific value exists in the linked list
func (dll *DoublyLinkedList[T]) Has(val T) bool {
	//	TODO:
	return false
}

// The Head method is used to obtain the value of the head node of the linked list
func (dll *DoublyLinkedList[T]) Head() (val T, ok bool) {
	if dll.Empty() {
		return val, false
	}
	return dll.head.Val, true
}

// The Tail method is used to obtain the value of the tail node of the linked list
func (dll *DoublyLinkedList[T]) Tail() (val T, ok bool) {
	if dll.Empty() {
		return val, false
	}
	return dll.tail.Val, true
}

// The Size method is used to retrieve the number of nodes in the linked list
func (dll *DoublyLinkedList[T]) Size() int {
	return dll.size
}

// The Empty method is used to check whether the linked list is empty
func (dll *DoublyLinkedList[T]) Empty() bool {
	return dll.size == 0
}

// The Slice method is used to convert the linked list into a slice
func (dll *DoublyLinkedList[T]) Slice() []T {
	slice := make([]T, 0, dll.size)
	head := dll.head
	for head != nil {
		slice = append(slice, head.Val)
	}
	return slice
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
