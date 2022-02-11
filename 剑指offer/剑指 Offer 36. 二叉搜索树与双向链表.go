package main

var first, last *TreeNode

func treeToDoublyList(root *TreeNode) *TreeNode {

	first, last = nil, nil
	if root == nil {
		return nil
	}
	treeToDoublyListDfs(root)
	last.Right = first
	first.Left = last
	return first
}

func treeToDoublyListDfs(node *TreeNode) {
	if node == nil {
		return
	}
	treeToDoublyListDfs(node.Left)
	if last == nil {
		first = node
	} else {
		last.Right = node
		node.Left = last
	}

	last = node

	treeToDoublyListDfs(node.Right)
}
