package main

// https://leetcode-cn.com/problems/jian-sheng-zi-lcof/
// dp[i]: 表示i被分成x个整数的最大乘积
// dp[i]: 可以被分成 j 和 i-j 两部分:
// 其中 i-j 可以继续再拆分 或者不拆分
// 因此 dp[i] = max(j * (i-j), j * dp[i-j])
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = max(curMax, max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = curMax
	}
	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
