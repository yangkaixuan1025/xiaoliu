package main

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)

	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := 1; i*i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j >= i*i {
				dp[j] = min(dp[j-i*i]+1, dp[j])
			}
		}
	}

	return dp[n]

}
