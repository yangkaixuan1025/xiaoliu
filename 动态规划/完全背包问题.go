package main

// test_CompletePack1 先遍历物品, 在遍历背包
func test_CompletePack1(weight, value []int, bagWeight int) int {
	// 定义dp数组 和初始化
	dp := make([]int, bagWeight+1)
	// 遍历顺序
	for i := 0; i < len(weight); i++ {
		// 正序会多次添加 value[i]
		for j := weight[i]; j <= bagWeight; j++ {
			// 推导公式
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
			// debug
			//fmt.Println(dp)
		}
	}
	return dp[bagWeight]
}

// test_CompletePack2 先遍历背包, 在遍历物品
func test_CompletePack2(weight, value []int, bagWeight int) int {
	// 定义dp数组 和初始化
	dp := make([]int, bagWeight+1)
	// 遍历顺序
	// j从0 开始
	for j := 0; j <= bagWeight; j++ {
		for i := 0; i < len(weight); i++ {
			if j >= weight[i] {
				// 推导公式
				dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
			}
			// debug
			//fmt.Println(dp)
		}
	}
	return dp[bagWeight]
}
func test_CompletePackTwo(weight, value []int, bagweight int) int {
	// 定义dp数组 和初始化
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
				if dp[i-1][j] > (dp[i][j-weight[i]] + value[i]) {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-weight[i]] + value[i]
				}
			}
		}
	}

	return dp[len(weight)-1][bagweight]
}
