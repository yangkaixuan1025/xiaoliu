package main

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < target || (sum+target)%2 == 1 || (sum+target) < 0 {
		return 0
	}
	//开始dp初始化
	weight := (target + sum) / 2
	dp := make([][]int, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]int, weight+1) //背包容量
	}
	dp[0][0] = 1 //当背包容量为0，且物品为0时，填满背包就1种方法
	for i := 1; i <= len(nums); i++ {

		for j := 0; j <= weight; j++ {
			if nums[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[len(nums)][weight]
}
