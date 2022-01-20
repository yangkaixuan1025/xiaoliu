package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	dummy := &ListNode{
		Next: head,
	}
	curr := dummy
	for curr.Next != nil && curr.Next.Next != nil {
		tmp := curr.Next.Next.Next
		tmp1 := curr.Next.Next
		tmp2 := curr.Next
		curr.Next = tmp1
		curr.Next.Next = tmp2
		curr.Next.Next.Next = tmp
		curr = tmp2
	}

	return dummy.Next

}

func swapPairsVersion2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}
