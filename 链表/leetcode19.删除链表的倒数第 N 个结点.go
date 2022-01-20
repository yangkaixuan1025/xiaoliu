package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 {
		return head
	}
	if head == nil {
		return head
	}
	dummy := &ListNode{
		Next: head,
	}
	curr := dummy
	for n > 0 && curr != nil {
		curr = curr.Next
		n--
	}
	if n > 0 {
		return head
	}
	if curr == nil {
		return head
	}
	pre := dummy
	for curr.Next != nil {
		pre = pre.Next
		curr = curr.Next
	}
	pre.Next = pre.Next.Next
	return dummy.Next
}
