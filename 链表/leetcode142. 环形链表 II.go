package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	nodeMap := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := nodeMap[head]; ok {
			return head
		} else {
			nodeMap[head] = struct{}{}
		}
		head = head.Next
	}
	return nil

}

func detectCycleVersion2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}
