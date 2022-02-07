package main

func numDistinct(s string, t string) int {
	dp := make([][]int, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(t))
	}
	// 初始化
	for i := 0; i < len(dp); i++ {
		dp[i][0] = 1
	}
	// dp[0][j] 为 0，默认值，因此不需要初始化
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
