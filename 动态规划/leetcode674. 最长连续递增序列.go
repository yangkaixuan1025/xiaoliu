package main

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	r := 1
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			dp[i+1] = dp[i] + 1
		}
		if dp[i+1] > r {
			r = dp[i+1]
		}
	}

	return r

}
