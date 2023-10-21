package structure

type Stacker[T comparable] interface {
	Push(val T)
	Pop() (val T, ok bool)
	Top() (val T, ok bool)
	Empty() bool
	Size() int
}

type Stack[T comparable] struct {
	ll *LinkedList[T]
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{
		ll: NewLinkedList[T](),
	}
}

func (s *Stack[T]) Push(val T) {
	s.ll.InsertHead(val)
}

func (s *Stack[T]) Pop() (val T, ok bool) {
	val, ok = s.ll.DeleteHead()
	return val, ok
}

func (s *Stack[T]) Top() (val T, ok bool) {
	return s.ll.Head()
}

func (s *Stack[T]) Empty() bool {
	return s.ll.Empty()
}

func (s *Stack[T]) Size() int {
	return s.ll.Size()
}

func (s *Stack[T]) Slice() []T {
	return s.ll.Slice()
}
