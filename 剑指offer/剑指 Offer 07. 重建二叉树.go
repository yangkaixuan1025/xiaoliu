package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序+中序  与  后序+中序 是可以重建二叉树的
// 前序+后序 没有办法唯一确定一个二叉树
// 前序第一个是root节点， 找到中序root节点， 那么root节点之前的是左子节点，后面是右子节点

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	v := preorder[0]
	root := &TreeNode{
		Val: v,
	}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == v {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(preorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(preorder[:i])+1:], inorder[i+1:])
	return root
}
