package main

import (
	"testing"
)

func TestInsertToHead(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToHead(i + 1)
	}
	l.Print()
}

func TestInsertToTail(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
}

func TestFindByIndex(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	t.Log(l.FindByIndex(0))
	t.Log(l.FindByIndex(9))
	t.Log(l.FindByIndex(5))
	t.Log(l.FindByIndex(11))
}

func TestDeleteNode(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 3; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()

	t.Log(l.DeleteNode(l.Head.Next))
	l.Print()

	t.Log(l.DeleteNode(l.Head.Next.Next))
	l.Print()
}

var l *LinkedList

func init() {
	n5 := &ListNode{Value: 5}
	n4 := &ListNode{Value: 4, Next: n5}
	n3 := &ListNode{Value: 3, Next: n4}
	n2 := &ListNode{Value: 2, Next: n3}
	n1 := &ListNode{Value: 1, Next: n2}
	l = &LinkedList{Head: &ListNode{Next: n1}}
}

func TestReverse(t *testing.T) {
	l.Print()
	l.Reverse()
	l.Print()
}

func TestHasCycle(t *testing.T) {
	t.Log(l.HasCycle())
	l.Head.Next.Next.Next.Next.Next.Next = l.Head.Next.Next.Next
	t.Log(l.HasCycle())
}

func TestMergeSortedList(t *testing.T) {
	n5 := &ListNode{Value: 9}
	n4 := &ListNode{Value: 7, Next: n5}
	n3 := &ListNode{Value: 5, Next: n4}
	n2 := &ListNode{Value: 3, Next: n3}
	n1 := &ListNode{Value: 1, Next: n2}
	l1 := &LinkedList{Head: &ListNode{Next: n1}}

	n10 := &ListNode{Value: 10}
	n9 := &ListNode{Value: 8, Next: n10}
	n8 := &ListNode{Value: 6, Next: n9}
	n7 := &ListNode{Value: 4, Next: n8}
	n6 := &ListNode{Value: 2, Next: n7}
	l2 := &LinkedList{Head: &ListNode{Next: n6}}

	MergeSortedList(l1, l2).Print()
}

func TestDeleteBottomN(t *testing.T) {
	l.Print()
	l.DeleteBottomN(3)
	l.Print()
}

func TestFindMiddleNode(t *testing.T) {
	//l.DeleteBottomN(1)
	//l.DeleteBottomN(1)
	//l.DeleteBottomN(1)
	//l.DeleteBottomN(1)
	l.Print()
	t.Log(l.FindMiddleNode())
}
