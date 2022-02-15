package main

func maxValue(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	if n == 0 {
		return 0
	}
	if m == 0 {
		return 0
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = max(dp[i-1][j]+grid[i][j], dp[i][j-1]+grid[i][j])
		}
	}
	return dp[n-1][m-1]

}
