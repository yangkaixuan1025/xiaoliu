package main

type MyLinkedList struct {
	dummy *Node
}

type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}

func Constructor() MyLinkedList {
	rear := &Node{
		Val: -1,
	}
	rear.Next = rear
	rear.Pre = rear
	return MyLinkedList{dummy: rear}
}

func (this *MyLinkedList) Get(index int) int {
	if this == nil || this.dummy == nil {
		return -1
	}
	curr := this.dummy.Next
	for index > 0 {
		if curr == this.dummy {
			return -1
		}
		index--
		curr = curr.Next
	}
	return curr.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	dummy := this.dummy
	node := &Node{
		Val:  val,
		Next: dummy.Next,
		Pre:  dummy,
	}

	dummy.Next.Pre = node
	dummy.Next = node
}

func (this *MyLinkedList) AddAtTail(val int) {
	dummy := this.dummy
	node := &Node{
		Val:  val,
		Next: dummy,
		Pre:  dummy.Pre,
	}
	dummy.Pre.Next = node
	dummy.Pre = node

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	curr := this.dummy.Next
	for index > 0 {
		if curr == this.dummy {
			return
		}
		index--
		curr = curr.Next
	}
	node := &Node{
		Val:  val,
		Next: curr,
		Pre:  curr.Pre,
	}
	curr.Pre.Next = node
	curr.Pre = node

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.dummy.Next == this.dummy {
		return
	}
	head := this.dummy.Next
	for head.Next != this.dummy && index > 0 {
		head = head.Next
		index--
	}
	if index == 0 {
		head.Next.Pre = head.Pre
		head.Pre.Next = head.Next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
