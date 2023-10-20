package xkit

// Iterator interface constraints
type Iterator[T any] interface {
	ForEach(func(item T))
}

// ToStringer interface constraints
type ToStringer interface {
	String() string
}

// ToSlicer interface constraints
type ToSlicer[T any] interface {
	Slice() []T
}

// ToMapper interface constraints
type ToMapper[K comparable, V any] interface {
	Map() map[K]V
}
