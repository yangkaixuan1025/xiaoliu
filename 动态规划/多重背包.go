package main

//有N种物品和一个容量为V 的背包。第i种物品最多有Mi件可用，每件耗费的空间是Ci ，价值是Wi 。求解将哪些物品装入背包可使这些物品的耗费的空间 总和不超过背包容量，且价值总和最大
func testMultiPack(weight, value, nums []int, bagWeight int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] > 1 {
			weight = append(weight, weight[i])
			value = append(value, value[i])
			nums[i]--
		}
	}
	dp := make([]int, bagWeight+1)
	for i := 0; i < len(weight); i++ {
		for j := bagWeight; j >= weight[i]; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}

	return dp[bagWeight]
}

func testMultiPack2(weight, value, nums []int, bagWeight int) int {
	dp := make([]int, bagWeight+1)
	for i := 0; i < len(weight); i++ {
		for j := bagWeight; j >= weight[i]; j-- {
			for k := 1; k <= nums[i] && j >= k*weight[i]; k++ {
				dp[j] = max(dp[j], dp[j-k*weight[i]]+k*value[i])
			}

		}
	}

	return dp[bagWeight]
}
