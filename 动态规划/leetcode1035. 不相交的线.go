package main

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	t1 := len(nums1)
	t2 := len(nums2)
	dp := make([][]int, t1+1)
	for i := range dp {
		dp[i] = make([]int, t2+1)
	}
	for i := 1; i <= t1; i++ {
		for j := 1; j <= t2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[t1][t2]
}
