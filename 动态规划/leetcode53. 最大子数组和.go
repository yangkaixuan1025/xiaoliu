package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, 1)
	dp[0] = nums[0]
	result := dp[0]
	for i := 1; i < len(nums); i++ {
		sum1 := dp[i-1] + nums[i]
		if sum1 > nums[i] {
			dp = append(dp, sum1)
		} else {
			dp = append(dp, nums[i])
		}
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}
