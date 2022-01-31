package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricDfs(root.Left, root.Right)

}

func isSymmetricDfs(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil && right != nil {
		return false
	}
	if left != nil && right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return isSymmetricDfs(left.Right, right.Left) && isSymmetricDfs(left.Left, right.Right)
}

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root.Left, root.Right)
	for len(queue) > 0 {
		left := queue[len(queue)-1]
		right := queue[len(queue)-2]
		queue = queue[:len(queue)-2]
		if left == nil && right == nil {
			continue
		}
		if left == nil && right != nil {
			return false
		}
		if left != nil && right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}
	return true

}
