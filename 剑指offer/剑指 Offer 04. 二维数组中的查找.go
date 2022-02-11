package main

//https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
//[
//[1,   4,  7, 11, 15],
//[2,   5,  8, 12, 19],
//[3,   6,  9, 16, 22],
//[10, 13, 14, 17, 24],
//[18, 21, 23, 26, 30]
//]

// 从左下角逐渐剔除行和列
// 如果大于目标值就剔除行
// 如果小于目标值就剔除列
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	n := len(matrix)
	m := len(matrix[0])
	i, j := n-1, 0
	for i >= 0 && j <= m-1 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			i--
			continue
		}
		if matrix[i][j] < target {
			j++
			continue
		}
	}
	return false

}
