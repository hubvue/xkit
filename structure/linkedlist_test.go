package structure

import "testing"

func testErrorf[T1 any, T2 any](t *testing.T, msg string, val T1, expected T2) {
	t.Errorf("%s = %v; expected = %v.", msg, val, expected)
}

func testError(t *testing.T, args ...any) {
	t.Error(args...)
}

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList[int]()
	if ll.size != 0 {
		testErrorf(t, "NewLinkedList().size", ll.size, 0)
	}
	if ll.head != nil {
		testErrorf(t, "NewLinkedList().head", ll.head, "nil")
	}
	ll.Add(1)
	if ll.size == 0 {
		testErrorf(t, "NewLinkedList().size", ll.size, 1)
	}
	if ll.head == nil {
		testErrorf(t, "NewLinkedList().head.Value", "nil", 1)
	}
}

func TestNewLinkedNode(t *testing.T) {
	n2 := NewLinkedNode[int](2, nil)
	n1 := NewLinkedNode[int](1, n2)
	if n1.Val != 1 {
		testErrorf(t, "NewLinkedNode().Val", n1.Val, 1)
	}
	if n1.Next == nil {
		testError(t, "LinkedNode().Next should a LinkedNode")
	}
}

func TestLinkedList_Add(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Add(1)
	ll.Add(2)
	if ll.size != 2 {
		testErrorf(t, "Two Add caller, ll.size", ll.size, 2)
	}
}

func TestLinkedList_Delete(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Add(1)
	ll.Add(2)
	ok := ll.Delete(1)
	if ok != true {
		testErrorf(t, "Delete(1), ok", ok, true)
	}
	if ll.size != 1 {
		testErrorf(t, "Delete(1), ll.size", ll.size, 1)
	}
	ok = ll.Delete(3)
	if ok != false {
		testErrorf(t, "Delete(3), ok", ok, false)
	}

}

func TestLinkedList_Insert(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Add(1)
	ok := ll.Insert(1, 2)
	if ll.size != 2 {
		testErrorf(t, "Insert(1, 2), ll.size", ll.size, 2)
	}
	if ok != true {
		testErrorf(t, "Insert(1, 2), ok", ok, true)
	}
	ok = ll.Insert(3, 4)
	if ok != false {
		testErrorf(t, "Insert(3, 4), ok", ok, false)
	}
	if ll.size != 2 {
		testErrorf(t, "Insert(1, 2), ll.size", ll.size, 2)
	}

}

func TestLinkedList_Empty(t *testing.T) {
	ll := NewLinkedList[int]()
	ok := ll.Empty()
	if ok != true {
		testErrorf(t, "Empty() is called after NewLinkedList(), ok", ok, true)
	}
	ll.Add(1)
	ok = ll.Empty()
	if ok != false {
		testErrorf(t, "Empty() is called after Add(1), ok", ok, false)
	}
}

func TestLinkedList_Has(t *testing.T) {
	ll := NewLinkedList[int]()
	ok := ll.Has(1)
	if ok != false {
		testError(t, "Has(1) is called after NewLinkedList(), ok", ok, false)
	}
	ll.Add(1)
	ll.Add(2)
	ok = ll.Has(2)
	if ok != true {
		testErrorf(t, "Has(1) is called after Add(1)", ok, true)
	}
}

func TestLinkedList_Head(t *testing.T) {
	ll := NewLinkedList[int]()
	_, ok := ll.Head()
	if ok != false {
		testErrorf(t, "Head() is called after NewLinkedList(), ok", ok, false)
	}
	ll.Add(1)
	val, ok := ll.Head()
	if ok != true {
		testErrorf(t, "Head() is called after Add(1), ok", ok, true)
	}
	if val != 1 {
		testErrorf(t, "Head() is called after Add(1), val", val, 1)
	}
	ll.Add(2)
	val, _ = ll.Head()
	if val != 1 {
		testErrorf(t, "Head() is called after Add(2), val", val, 1)
	}

}

func TestLinkedList_Reverse(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Reverse()
	val, _ := ll.Head()
	if val != 3 {
		testErrorf(t, "Reverse(), val", val, 3)
	}
}

func TestLinkedList_Size(t *testing.T) {
	ll := NewLinkedList[int]()
	size := ll.Size()
	if size != 0 {
		testErrorf(t, "Size()", size, 0)
	}
	ll.Add(1)
	size = ll.Size()
	if size != 1 {
		testErrorf(t, "Size() is called after Add(1)", size, 1)
	}
	ll.Insert(1, 2)
	size = ll.Size()
	if size != 2 {
		testErrorf(t, "Size() is called after Insert(1, 2)", size, 2)
	}
	ll.Delete(2)
	size = ll.Size()
	if size != 1 {
		testErrorf(t, "Size() is called after Delete(2)", size, 1)
	}
}

func TestLinkedList_Update(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Add(1)
	ll.Add(2)
	ok := ll.Update(1, 3)
	if ok != true {
		testErrorf(t, "Update(1, 3), ok", ok, true)
	}
	ok = ll.Update(4, 5)
	if ok != false {
		testErrorf(t, "Update(4, 5), ok", ok, false)
	}
	val, _ := ll.Head()
	if val != 3 {
		testErrorf(t, "Update(1, 3), val", val, 3)
	}
}

func TestLinkedList_ToString(t *testing.T) {
	ll := NewLinkedList[int]()
	s := ll.ToString()
	if s != "nil" {
		testErrorf(t, "ToString(), s", s, "nil")
	}
	ll.Add(1)
	ll.Add(2)
	s = ll.ToString()
	expected := "1 -> 2 -> nil"
	if s != expected {
		testErrorf(t, "ToString(), s", s, expected)
	}
}
