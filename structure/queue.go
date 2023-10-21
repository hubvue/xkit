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
	front *LinkedNode[T]
	rear  *LinkedNode[T]
	size  int
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(val T) {
	newNode := NewLinkedNode(val, nil)
	if q.front == nil {
		q.front = newNode
	} else {
		q.rear.Next = newNode
	}
	q.rear = newNode
	q.size++
}

func (q *Queue[T]) Dequeue() (val T, ok bool) {
	if q.Empty() {
		return val, false
	}
	val = q.rear.Val
	q.rear = q.rear.Next
	if q.front == q.rear {
		q.front = q.front.Next
	}
	q.size--
	return val, true
}

func (q *Queue[T]) Front() (val T, ok bool) {
	if q.Empty() {
		return val, false
	}
	return q.front.Val, true
}

func (q *Queue[T]) Rear() (val T, ok bool) {
	if q.Empty() {
		return val, false
	}
	return q.rear.Val, true
}

func (q *Queue[T]) Empty() bool {
	return q.size == 0
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) Slice() []T {
	//	TODO:
	return make([]T, 0)
}
