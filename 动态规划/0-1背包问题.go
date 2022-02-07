package main

//有n件物品和一个最多能背重量为w 的背包。第i件物品的重量是weight[i]，得到的价值是value[i] 。每件物品只能用一次，求解将哪些物品装入背包里物品价值总和最大。
func test_2_wei_bag_problem1(weight, value []int, bagweight int) int {
	dp := make([][]int, len(weight))
	for i := range dp {
		dp[i] = make([]int, bagweight+1)
	}
	for j := weight[0]; j <= bagweight; j++ {
		dp[0][j] = value[0]
	}
	for i := 0; i < len(weight); i++ {
		for j := 0; j <= bagweight; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				if dp[i-1][j] > (dp[i-1][j-weight[i]] + value[i]) {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i-1][j-weight[i]] + value[i]
				}
			}
		}
	}

	return dp[len(weight)-1][bagweight]
}

func test_1_wei_bag_problem(weight, value []int, bagWeight int) int {
	dp := make([]int, bagWeight+1)
	for i := 0; i < len(weight); i++ {
		for j := bagWeight; j >= weight[j]; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]+value[i]])
		}

	}
	return dp[bagWeight]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
