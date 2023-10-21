package structure

type CircularQueueable[T comparable] interface {
	Enqueue(val T) bool
	Dequeue() (val T, ok bool)
	Front() (val T, ok bool)
	Rear() (val T, ok bool)
	Empty() bool
	Size() int
	Full() bool
}

type CircularQueue[T comparable] struct {
	cap   int
	front int
	rear  int
	size  int
	raw   []T
}

func NewCircularQueue[T comparable](cap int) *CircularQueue[T] {
	raw := make([]T, cap)
	return &CircularQueue[T]{
		cap: cap,
		raw: raw,
	}
}

func (q *CircularQueue[T]) Enqueue(val T) bool {
	if q.Full() {
		return false
	}
	q.raw[q.rear] = val
	q.rear = (q.rear + 1) % q.cap
	q.size++
	return true
}

func (q *CircularQueue[T]) Dequeue() (val T, ok bool) {
	if q.Empty() {
		return val, false
	}
	val = q.raw[q.front]
	q.front = (q.front + 1) % q.cap
	q.size--
	return val, true
}

func (q *CircularQueue[T]) Front() (val T, ok bool) {
	if q.Empty() {
		return val, false
	}
	return q.raw[q.front], true
}

func (q *CircularQueue[T]) Rear() (val T, ok bool) {
	if q.Empty() {
		return val, false
	}
	rearIdx := ((q.rear - 1) + q.cap) % q.cap
	return q.raw[rearIdx], true
}

func (q *CircularQueue[T]) Empty() bool {
	return q.size == 0 && q.front == q.rear
}

func (q *CircularQueue[T]) Full() bool {
	return q.size == q.cap && q.front == q.rear
}

func (q *CircularQueue[T]) Size() int {
	return q.size
}

func (q *CircularQueue[T]) Slice() []T {
	queue := make([]T, q.size)
	if q.Empty() {
		return queue
	}
	queue[0] = q.raw[q.front]
	idx := q.front + 1
	i := 1
	for {
		if idx == q.rear {
			break
		}
		queue[i] = q.raw[idx]
		i++
		idx = (idx + 1) % q.cap
	}
	return queue
}

func (q *CircularQueue[T]) ForEach(handler func(item T)) {
	if q.Empty() {
		return
	}
	handler(q.raw[q.front])
	idx := q.front + 1
	for {
		if idx == q.rear {
			break
		}
		handler(q.raw[idx])
		idx = (idx + 1) % q.cap
	}
}
