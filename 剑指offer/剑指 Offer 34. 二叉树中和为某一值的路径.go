package main

var res [][]int

func pathSum(root *TreeNode, target int) [][]int {
	res = make([][]int, 0)
	if root == nil {
		return res
	}
	pathSumDfs(root, 0, target, []int{})
	return res

}

func pathSumDfs(root *TreeNode, sum, target int, path []int) {
	if root == nil {
		return
	}
	sum += root.Val
	path = append(path, root.Val)
	defer func() {
		path = path[:len(path)-1]

	}()
	if root.Left == nil && root.Right == nil && sum == target {
		res = append(res, append([]int(nil), path...))
		return
	}
	pathSumDfs(root.Left, sum, target, path)
	pathSumDfs(root.Right, sum, target, path)
}
