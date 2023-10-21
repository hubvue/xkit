package structure

import (
	"testing"
)

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
	ll.InsertTail(1)
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

func TestLinkedList_InsertTail(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.InsertTail(1)
	ll.InsertTail(2)
	if ll.size != 2 {
		testErrorf(t, "Two InsertTail caller, ll.size", ll.size, 2)
	}
}

func TestLinkedList_InsertHead(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.InsertHead(1)
	if ll.head == nil {
		testErrorf(t, "After the method call InsertHead(1), ll.head", ll.head, "not nil")
	}
	if ll.head.Val != 1 {
		testErrorf(t, "After the method call InsertHead(1), ll.head.Val", ll.head.Val, 1)
	}
	ll.InsertHead(2)
	if ll.head.Val != 2 {
		testErrorf(t, "After the method call InsertHead(2), ll.head.Val", ll.head.Val, 2)
	}
}

func TestLinkedList_InsertBefore(t *testing.T) {
	ll := NewLinkedList[int]()
	ok := ll.InsertBefore(1, 2)
	if ok != false {
		testErrorf(t, "InsertBefore(1, 2), ok", ok, false)
	}
	ll.InsertTail(1)
	if ll.head == nil {
		testErrorf(t, "After the method call InsertTail(1), ll.head", ll.head, "not nil")
	}
	if ll.head.Val != 1 {
		testErrorf(t, "After the method call InsertTail(1), ll.head.Val", ll.head.Val, 1)
	}
	ok = ll.InsertBefore(1, 2)
	if ok != true {
		testErrorf(t, "InsertBefore(1, 2), ok", ok, false)
	}
	if ll.head.Val != 2 {
		testErrorf(t, "After the method call InsertTail(2), ll.head.Val", ll.head.Val, 2)
	}
	if ll.head.Next == nil {
		testErrorf(t, "After the method call InsertTail(2), ll.head.Next", ll.head.Next, "not nil")
	}
	if ll.head.Next.Val != 1 {
		testErrorf(t, "After the method call InsertTail(2), ll.head.Next.Val", ll.head.Next.Val, 1)
	}
	ok = ll.InsertBefore(1, 3)
	if ok != true {
		testErrorf(t, "InsertBefore(1, 2), ok", ok, false)
	}
	val, _ := ll.Tail()
	if val != 1 {
		testErrorf(t, "InsertBefore(1, 3), Tail().val", val, 1)
	}
	ok = ll.InsertBefore(4, 5)
	if ok != false {
		testErrorf(t, "InsertBefore(4, 5), ok", ok, false)
	}
}

func TestLinkedList_InsertAfter(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.InsertTail(1)
	ok := ll.InsertAfter(1, 2)
	if ll.size != 2 {
		testErrorf(t, "InsertAfter(1, 2), ll.size", ll.size, 2)
	}
	if ok != true {
		testErrorf(t, "InsertAfter(1, 2), ok", ok, true)
	}
	ok = ll.InsertAfter(3, 4)
	if ok != false {
		testErrorf(t, "InsertAfter(3, 4), ok", ok, false)
	}
	if ll.size != 2 {
		testErrorf(t, "InsertAfter(1, 2), ll.size", ll.size, 2)
	}

}

func TestLinkedList_Update(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.InsertTail(1)
	ll.InsertTail(2)
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

func TestLinkedList_DeleteHead(t *testing.T) {
	ll := NewLinkedList[int]()
	_, ok := ll.DeleteHead()
	if ok != false {
		testErrorf(t, "The first call to the DeleteHead method after creating an empty linked list, ok", ok, false)
	}
	ll.InsertHead(1)
	val, ok := ll.DeleteHead()
	if ok != true {
		testErrorf(t, "Calling the DeleteHead method after InsertHead, ok", ok, true)
	}
	if val != 1 {
		testErrorf(t, "Calling the DeleteHead(1) method after InsertHead, val", val, true)
	}
	_, ok = ll.DeleteHead()
	if ok != false {
		testErrorf(t, "DeleteTail, ok", ok, false)
	}

}

func TestLinkedList_DeleteTail(t *testing.T) {
	ll := NewLinkedList[int]()
	_, ok := ll.DeleteTail()
	if ok != false {
		testErrorf(t, "The first call to the DeleteTail method after creating an empty linked list, ok", ok, false)
	}
	ll.InsertHead(1)
	ll.InsertHead(2)
	val, ok := ll.DeleteTail()
	if ok != true {
		testErrorf(t, "Calling the DeleteTail method after InsertHead, ok", ok, true)
	}
	if val != 1 {
		testErrorf(t, "insert: InsertHead(1), InsertHead(2), val", val, 1)
	}
	ll.InsertHead(3)
	ll.InsertHead(4)
	ll.InsertHead(5)
	val, ok = ll.DeleteTail()
	if ok != true {
		testErrorf(t, "DeleteTail(), ok", ok, true)
	}
	if val != 2 {
		testErrorf(t, "DeleteTail(), val", val, 2)
	}
	_, ok = ll.DeleteTail()
	_, ok = ll.DeleteTail()
	_, ok = ll.DeleteTail()
	_, ok = ll.DeleteTail()
	if ok != false {
		testErrorf(t, "DeleteTail(), ok", ok, false)
	}
}

func TestLinkedList_Delete(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.InsertTail(1)
	ll.InsertTail(2)
	ok := ll.Delete(1)
	if ok != true {
		testErrorf(t, "Delete(1), ok", ok, true)
	}
	if ll.size != 1 {
		testErrorf(t, "Delete(1), ll.size", ll.size, 1)
	}
	ll.InsertTail(3)
	ll.InsertTail(4)
	ll.InsertTail(5)
	ll.InsertTail(6)
	ll.InsertTail(7)
	ok = ll.Delete(8)
	if ok != false {
		testErrorf(t, "Delete(3), ok", ok, false)
	}

}

func TestLinkedList_Empty(t *testing.T) {
	ll := NewLinkedList[int]()
	ok := ll.Empty()
	if ok != true {
		testErrorf(t, "Empty() is called after NewLinkedList(), ok", ok, true)
	}
	ll.InsertTail(1)
	ok = ll.Empty()
	if ok != false {
		testErrorf(t, "Empty() is called after InsertTail(1), ok", ok, false)
	}
}

func TestLinkedList_Has(t *testing.T) {
	ll := NewLinkedList[int]()
	ok := ll.Has(1)
	if ok != false {
		testError(t, "Has(1) is called after NewLinkedList(), ok", ok, false)
	}
	ll.InsertTail(1)
	ll.InsertTail(2)
	ok = ll.Has(2)
	if ok != true {
		testErrorf(t, "Has(1) is called after InsertTail(1)", ok, true)
	}
}

func TestLinkedList_Head(t *testing.T) {
	ll := NewLinkedList[int]()
	_, ok := ll.Head()
	if ok != false {
		testErrorf(t, "Head() is called after NewLinkedList(), ok", ok, false)
	}
	ll.InsertTail(1)
	val, ok := ll.Head()
	if ok != true {
		testErrorf(t, "Head() is called after InsertTail(1), ok", ok, true)
	}
	if val != 1 {
		testErrorf(t, "Head() is called after InsertTail(1), val", val, 1)
	}
	ll.InsertTail(2)
	val, _ = ll.Head()
	if val != 1 {
		testErrorf(t, "Head() is called after InsertTail(2), val", val, 1)
	}

}

func TestLinkedList_Tail(t *testing.T) {
	ll := NewLinkedList[int]()
	_, ok := ll.Tail()
	if ok != false {
		testErrorf(t, "Tail(), ok", ok, false)
	}
	ll.InsertTail(1)
	val, ok := ll.Tail()
	if ok != true {
		testErrorf(t, "Tail(), ok", ok, true)
	}
	if val != 1 {
		testErrorf(t, "Tail(), val", val, 1)
	}
	ll.InsertTail(2)
	val, ok = ll.Tail()
	if ok != true {
		testErrorf(t, "Tail(), ok", ok, true)
	}
	if val != 2 {
		testErrorf(t, "Tail(), val", val, 2)
	}

}

func TestLinkedList_Reverse(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.InsertTail(1)
	ll.InsertTail(2)
	ll.InsertTail(3)
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
	ll.InsertTail(1)
	size = ll.Size()
	if size != 1 {
		testErrorf(t, "Size() is called after InsertTail(1)", size, 1)
	}
	ll.InsertAfter(1, 2)
	size = ll.Size()
	if size != 2 {
		testErrorf(t, "Size() is called after InsertAfter(1, 2)", size, 2)
	}
	ll.Delete(2)
	size = ll.Size()
	if size != 1 {
		testErrorf(t, "Size() is called after Delete(2)", size, 1)
	}
	ll.InsertHead(1)
	if ll.Size() != 2 {
		testErrorf(t, "Size(), size", ll.Size(), 2)
	}
	ll.DeleteHead()
	if ll.Size() != 1 {
		testErrorf(t, "Size(), size", ll.Size(), 1)
	}
	ll.DeleteTail()
	if ll.Size() != 0 {
		testErrorf(t, "Size(), size", ll.Size(), 0)
	}
}

func TestLinkedList_String(t *testing.T) {
	ll := NewLinkedList[int]()
	s := ll.String()
	if s != "nil" {
		testErrorf(t, "ToString(), s", s, "nil")
	}
	ll.InsertTail(1)
	ll.InsertTail(2)
	s = ll.String()
	expected := "1 -> 2 -> nil"
	if s != expected {
		testErrorf(t, "ToString(), s", s, expected)
	}
}

func TestLinkedList_Slice(t *testing.T) {
	ll := NewLinkedList[int]()
	s := ll.Slice()
	if len(s) != 0 {
		testErrorf(t, "Slice(), len", len(s), 0)
	}
	ll.InsertTail(0)
	ll.InsertTail(1)
	ll.InsertTail(2)
	s = ll.Slice()
	if len(s) != ll.Size() {
		testErrorf(t, "Slice(), len", len(s), ll.Size())
	}
}

func testTypeHandler[T comparable](t *testing.T, expected T) {
	ll := NewLinkedList[T]()
	ll.InsertHead(expected)
	val, _ := ll.Head()
	if val != expected {
		testErrorf(t, "Head(), val", val, expected)
	}
}
func TestLinkedList_Type(t *testing.T) {
	testTypeHandler[float32](t, 1.2)
	testTypeHandler[float64](t, 1.23)
	testTypeHandler[int](t, 1)
	testTypeHandler[int16](t, 1111)
	testTypeHandler[int32](t, 111111)
	testTypeHandler[int64](t, 11111)
	testTypeHandler[string](t, "test")
}
