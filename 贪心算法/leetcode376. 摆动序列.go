package main

func wiggleMaxLength(nums []int) int {
	count := 0
	if len(nums) <= 1 {
		return 1
	}
	preDiff := 0
	currDiff := 0
	count = 1
	for i := 0; i < len(nums)-1; i++ {
		currDiff = nums[i+1] - nums[i]
		if (preDiff >= 0 && currDiff < 0) || (preDiff <= 0 && currDiff > 0) {
			count++
			preDiff = currDiff
		}
	}
	return count
}
