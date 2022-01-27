package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	if len(stack) != 0 {
		e := stack[len(stack)-1]
		res = append(res, e.Val)
		stack = stack[:len(stack)-1]
		if e.Right != nil {
			stack = append(stack, e.Right)
		}
		if e.Left != nil {
			stack = append(stack, e.Left)
		}
	}
	return res
}

func preOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		if root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return res
}

func inOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := inOrderTraversal(root.Left)
	res = append(res, root.Val)
	res = append(res, inOrderTraversal(root.Right)...)
	return res
}

func inOrder(root *TreeNode) []int {
	res := make([]int, 0)
	if root != nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return res
}

func postOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	if root.Left != nil {
		res = append(res, postOrderTraversal(root.Left)...)

	}
	if root.Right != nil {
		res = append(res, postOrderTraversal(root.Right)...)

	}
	res = append(res, root.Val)
	return res
}
func postOrder(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	pre := &TreeNode{}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		if (cur.Right == nil && cur.Left == nil) || (pre != nil && (pre == cur.Left || pre == cur.Right)) {
			pre = cur
			res = append(res, cur.Val)
			stack = stack[:len(stack)-1]
		} else {
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
		}

	}
	return res
}

// 双栈实现后序遍历
//func (this *BinaryTree) PostOrderTraverse() {
//	s1 := NewArrayStack()
//	s2 := NewArrayStack()
//	s1.Push(this.root)
//	for !s1.IsEmpty() {
//		p := s1.Pop().(*Node)
//		s2.Push(p)
//		if nil != p.left {
//			s1.Push(p.left)
//		}
//		if nil != p.right {
//			s1.Push(p.right)
//		}
//	}
//
//	for !s2.IsEmpty() {
//		fmt.Printf("%+v ", s2.Pop().(*Node).data)
//	}
//}

// 层序遍历递归
var res [][]int

func levelOrder(root *TreeNode) [][]int {
	res = [][]int{}
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, level int) {
	if root != nil {
		if len(res) == level {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
}

// 层序遍历迭代

func levelOrder2(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return nil
	}
	var queue = []*TreeNode{root}
	var level int
	for len(queue) > 0 {
		counter := len(queue)
		res = append(res, []int{})
		for counter > 0 {
			counter--
			q := queue[0]

			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
			res[level] = append(res[level], q.Val)
			queue = queue[1:]
		}
		level++
	}
	return res
}
