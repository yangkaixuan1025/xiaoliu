package main

import "math"

func maxSubArray(nums []int) int {
	sum := int(math.MinInt64)
	if len(nums) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count >= sum {
			sum = count
		}
		if count < 0 {
			count = 0
		}
	}
	return sum
}

func maxSubArrayDP(nums []int) int {
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
