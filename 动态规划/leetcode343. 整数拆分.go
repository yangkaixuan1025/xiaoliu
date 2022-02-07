package main

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		r := 0
		for j := 1; j < i-1; j++ {
			if (j * (i - j)) > dp[i-j]*j {
				r = j * (i - j)
			} else {
				r = dp[i-j] * j
			}
			if r > dp[i] {
				dp[i] = r
			}
		}
	}
	return dp[n]
}
