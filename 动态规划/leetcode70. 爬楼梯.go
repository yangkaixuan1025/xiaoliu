package main

func climbStairs(n int) int {
	r := 0
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	pre := 2
	prePre := 1
	for i := 3; i <= n; i++ {
		r = pre + prePre
		prePre = pre
		pre = r
	}

	return r
}

// 完全背包问题来解决
func climbStairs2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= 2; j++ {
			if i >= j {
				dp[i] += dp[i-j]
			}
		}

	}
	return dp[n]
}
