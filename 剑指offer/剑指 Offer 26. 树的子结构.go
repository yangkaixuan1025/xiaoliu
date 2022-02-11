package main

// https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return A != nil && B != nil && (isSubStructureDfs(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))

}

func isSubStructureDfs(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}

	return isSubStructureDfs(A.Left, B.Left) && isSubStructureDfs(A.Right, B.Right)

}
