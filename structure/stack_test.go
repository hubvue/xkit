package structure

import "testing"

func TestNewStack(t *testing.T) {
	s := NewStack[int]()
	if s == nil {
		testErrorf(t, "NewStack(), s", s, "not nil")
	}
	if s.ll == nil {
		testErrorf(t, "NewStack(), s.ll", s.ll, "not nil")
	}
}

func TestStack_Push(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	if s.ll.Size() != 2 {
		testErrorf(t, "Two Push caller, size", s.Size(), 2)
	}
}

func TestStack_Pop(t *testing.T) {
	s := NewStack[int]()
	_, ok := s.Pop()
	if ok != false {
		testErrorf(t, "Pop(), ok", ok, false)
	}
	s.Push(1)
	s.Push(2)
	val, ok := s.Pop()
	if ok != true {
		testErrorf(t, "Pop(), ok", ok, true)
	}
	if val != 2 {
		testErrorf(t, "Pop(), val", val, 2)
	}
	val, ok = s.Pop()
	if ok != true {
		testErrorf(t, "Pop(), ok", ok, true)
	}
	if val != 1 {
		testErrorf(t, "Pop(), val", val, 1)
	}
	_, ok = s.Pop()
	if ok != false {
		testErrorf(t, "Pop(), ok", ok, false)
	}

}

func TestStack_Top(t *testing.T) {
	s := NewStack[int]()
	_, ok := s.Top()
	if ok != false {
		testErrorf(t, "Top(), ok", ok, false)
	}
	s.Push(1)
	val, ok := s.Top()
	if ok != true {
		testErrorf(t, "Top(), ok", ok, true)
	}
	if val != 1 {
		testErrorf(t, "Top(), val", val, 1)
	}
	s.Push(2)
	val, ok = s.Top()
	if ok != true {
		testErrorf(t, "Top(), ok", ok, true)
	}
	if val != 2 {
		testErrorf(t, "Top(), val", ok, 2)
	}
	s.Pop()
	val, ok = s.Top()
	if ok != true {
		testErrorf(t, "Top(), ok", ok, true)
	}
	if val != 1 {
		testErrorf(t, "Top(), val", val, 1)
	}
}

func TestStack_Size(t *testing.T) {
	s := NewStack[int]()
	if s.Size() != 0 {
		testErrorf(t, "Size(), size", s.Size(), 0)
	}
	s.Push(1)
	s.Push(2)
	if s.Size() != 2 {
		testErrorf(t, "Size(), size", s.Size(), 2)
	}
	s.Pop()
	if s.Size() != 1 {
		testErrorf(t, "Size(), size", s.Size(), 1)
	}
}

func TestStack_Empty(t *testing.T) {
	s := NewStack[int]()
	if !s.Empty() {
		testErrorf(t, "Empty(), empty", s.Empty(), true)
	}
	s.Push(1)
	if s.Empty() {
		testErrorf(t, "Empty(), empty", s.Empty(), false)
	}
	s.Pop()
	if !s.Empty() {
		testErrorf(t, "Empty(), empty", s.Empty(), true)
	}
}

func TestStack_Slice(t *testing.T) {
	s := NewStack[int]()
	ss := s.Slice()
	if len(ss) != s.Size() {
		testErrorf(t, "Slice(), size", len(ss), s.Size())
	}
	s.Push(1)
	s.Push(1)
	s.Push(1)
	s.Push(1)
	s.Push(1)
	ss = s.Slice()
	if len(ss) != s.Size() {
		testErrorf(t, "Slice(), size", len(ss), s.Size())
	}
}
