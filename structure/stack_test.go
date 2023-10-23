package structure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack[int]()
	assert.NotNil(t, s)
	assert.NotNil(t, s.ll)
}

func TestStack_Push(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	assert.Equal(t, s.ll.Size(), 2)
}

func TestStack_Pop(t *testing.T) {
	s := NewStack[int]()
	_, ok := s.Pop()
	assert.Equal(t, ok, false)
	s.Push(1)
	s.Push(2)
	val, ok := s.Pop()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 2)
	val, ok = s.Pop()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 1)
	_, ok = s.Pop()
	assert.Equal(t, ok, false)
}

func TestStack_Top(t *testing.T) {
	s := NewStack[int]()
	_, ok := s.Top()
	assert.Equal(t, ok, false)
	s.Push(1)
	val, ok := s.Top()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 1)
	s.Push(2)
	val, ok = s.Top()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 2)
	s.Pop()
	val, ok = s.Top()
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 1)
}

func TestStack_Size(t *testing.T) {
	s := NewStack[int]()
	assert.Equal(t, s.Size(), 0)
	s.Push(1)
	s.Push(2)
	assert.Equal(t, s.Size(), 2)
	s.Pop()
	assert.Equal(t, s.Size(), 1)
}

func TestStack_Empty(t *testing.T) {
	s := NewStack[int]()
	assert.Equal(t, s.Empty(), true)
	s.Push(1)
	assert.Equal(t, s.Empty(), false)
	s.Pop()
	assert.Equal(t, s.Empty(), true)
}

func TestStack_Slice(t *testing.T) {
	s := NewStack[int]()
	ss := s.Slice()
	assert.Equal(t, len(ss), s.Size())
	s.Push(1)
	s.Push(1)
	s.Push(1)
	s.Push(1)
	s.Push(1)
	ss = s.Slice()
	assert.Equal(t, len(ss), s.Size())
}
