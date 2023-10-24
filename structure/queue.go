package structure

type Queueable[T comparable] interface {
	Enqueue(val T)
	Dequeue() (val T, ok bool)
	Front() (val T, ok bool)
	Rear() (val T, ok bool)
	Empty() bool
	Size() int
}

type Queue[T comparable] struct {
	dll *DoublyLinkedList[T]
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{
		dll: NewDoublyLinkedList[T](),
	}
}

func (q *Queue[T]) Enqueue(val T) {
	q.dll.InsertTail(val)
}

func (q *Queue[T]) Dequeue() (val T, ok bool) {
	return q.dll.DeleteHead()
}

func (q *Queue[T]) Front() (val T, ok bool) {
	return q.dll.Tail()
}

func (q *Queue[T]) Rear() (val T, ok bool) {
	return q.dll.Head()
}

func (q *Queue[T]) Empty() bool {
	return q.dll.Empty()
}

func (q *Queue[T]) Size() int {
	return q.dll.Size()
}

func (q *Queue[T]) Slice() []T {
	return q.dll.Slice()
}
