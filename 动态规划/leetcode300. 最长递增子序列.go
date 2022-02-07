package main

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	r := 0
	for i := 1; i < len(nums); i++ {

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)

			}

		}
		if r < dp[i] {
			r = dp[i]

		}
	}
	return r
}
