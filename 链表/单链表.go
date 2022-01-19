package main

import "fmt"

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

type LinkedList struct {
	Head   *ListNode
	Length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{
		Value: v,
		Next:  nil,
	}
}

func (this *ListNode) GetNext() *ListNode {
	return this.Next
}

func (this *ListNode) GetValue() interface{} {
	return this.Value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head:   NewListNode(0),
		Length: 0,
	}
}

//在某个节点后面插入节点
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if p == nil {
		return false
	}
	newNode := NewListNode(v)
	newNode.Next = p.Next
	p.Next = newNode
	this.Length++
	return true
}

//在某个节点前面插入节点
func (this *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if p == nil || this.Head == p {
		return false
	}
	pre := this.Head
	cur := this.Head.Next
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.Next
	}
	if cur == nil {
		return false
	}
	newNode := NewListNode(v)
	pre.Next = newNode
	newNode.Next = cur
	this.Length++
	return true
}

//在链表头部插入节点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.Head, v)
}

//在链表尾部插入节点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.Head
	for nil != cur.Next {
		cur = cur.Next
	}
	return this.InsertAfter(cur, v)
}

//通过索引查找节点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= this.Length {
		return nil
	}
	cur := this.Head.Next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.Next
	}
	return cur
}

//删除传入的节点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if p == nil {
		return false
	}
	cur := this.Head.Next
	pre := this.Head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.Next
	}
	if cur == nil {
		return false
	}
	pre.Next = p.Next
	p = nil
	this.Length--
	return true
}

//打印链表
func (this *LinkedList) Print() {
	cur := this.Head.Next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.Next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

/*
单链表反转
时间复杂度：O(N)
*/
func (this *LinkedList) Reverse() {
	if nil == this.Head || nil == this.Head.Next || nil == this.Head.Next.Next {
		return
	}
	var pre *ListNode = nil
	cur := this.Head.Next
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp

		// 平行赋值，原子操作
		// cur.Next, pre, cur = pre, cur, cur.Next

	}
	this.Head.Next = pre

}

/*
判断单链表是否有环
*/
func (this *LinkedList) HasCycle() bool {
	if nil != this.Head {
		slow := this.Head
		fast := this.Head
		for nil != fast && nil != fast.Next {
			slow = slow.Next
			fast = fast.Next.Next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

/*
两个有序单链表合并
*/
func MergeSortedList(l1, l2 *LinkedList) *LinkedList {
	if l1 == nil || l1.Head == nil || l1.Head.Next == nil {
		return l2
	}
	if l2 == nil || l2.Head == nil || l2.Head.Next == nil {
		return l1
	}
	l := NewLinkedList()
	cur := l.Head
	cur1 := l1.Head.Next
	cur2 := l2.Head.Next
	for cur1 != nil && cur2 != nil {
		if cur1.Value.(int) > cur2.Value.(int) {
			cur.Next = cur2
			cur2 = cur2.Next
		} else {
			cur.Next = cur1
			cur1 = cur1.Next
		}
		cur = cur.Next
	}
	if cur1 != nil {
		cur.Next = cur1
	} else {
		cur.Next = cur2
	}
	return l
}

/*
删除倒数第N个节点
*/
func (this *LinkedList) DeleteBottomN(n int) {
	if this == nil || this.Head == nil || this.Head.Next == nil || n < 0 {
		return
	}
	fast := this.Head
	for i := 0; i < n && fast != nil; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return
	}

	slow := this.Head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next

}

/*
获取中间节点
*/
func (this *LinkedList) FindMiddleNode() *ListNode {
	if this == nil || this.Head == nil || this.Head.Next == nil {
		return nil
	}
	if this.Head.Next.Next == nil {
		return this.Head.Next
	}
	slow, fast := this.Head, this.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
