package structure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()
	assert.NotNil(t, q)
	assert.NotNil(t, q.dll)
}

func TestQueue_Dequeue(t *testing.T) {
	q := NewQueue[int]()
	_, ok := q.Dequeue()
	assert.Equal(t, ok, false)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	val, ok := q.Dequeue()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 1)
	val, ok = q.Dequeue()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 2)
	val, ok = q.Dequeue()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 3)
	val, ok = q.Dequeue()
	assert.Equal(t, ok, false)
}

func TestQueue_Enqueue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.Equal(t, q.dll.Size(), 3)
}

func TestQueue_Front(t *testing.T) {
	q := NewQueue[int]()
	_, ok := q.Front()
	assert.Equal(t, ok, false)
	q.Enqueue(1)
	val, ok := q.Front()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 1)
	q.Enqueue(2)
	val, ok = q.Front()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 2)
}

func TestQueue_Rear(t *testing.T) {
	q := NewQueue[int]()
	_, ok := q.Rear()
	assert.Equal(t, ok, false)
	q.Enqueue(1)
	val, ok := q.Rear()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 1)
	q.Enqueue(2)
	val, ok = q.Rear()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 1)
	q.Dequeue()
	val, ok = q.Rear()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 2)
}

func TestQueue_Empty(t *testing.T) {
	q := NewQueue[int]()
	assert.Equal(t, q.Empty(), true)
	q.Enqueue(1)
	assert.Equal(t, q.Empty(), false)
	q.Dequeue()
	assert.Equal(t, q.Empty(), true)
}

func TestQueue_Size(t *testing.T) {
	q := NewQueue[int]()
	assert.Equal(t, q.Size(), 0)
	q.Enqueue(1)
	assert.Equal(t, q.Size(), 1)
	q.Dequeue()
	assert.Equal(t, q.Size(), 0)
}

func TestQueue_Slice(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	s := q.Slice()
	for _, v := range s {
		val, _ := q.Dequeue()
		assert.Equal(t, val, v)
	}

}
