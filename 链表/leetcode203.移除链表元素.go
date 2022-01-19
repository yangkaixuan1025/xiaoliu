package main

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	newHead := head
	for newHead != nil {
		if newHead.Val == val {
			newHead = newHead.Next
		} else {
			break
		}
	}
	if newHead == nil {
		return nil
	}
	pre := newHead
	curr := newHead.Next
	for curr != nil {
		if curr.Val == val {
			pre.Next = curr.Next
		} else {
			pre = pre.Next
		}
		curr = curr.Next
	}
	return newHead
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 使用虚拟头结点
func removeElementsVersion2(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
