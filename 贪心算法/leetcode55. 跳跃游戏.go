package main

func canJump(nums []int) bool {
	maxIndex := 0
	for i := 0; i <= maxIndex; i++ {
		if nums[i]+i > maxIndex {
			maxIndex = nums[i] + i
		}
		if maxIndex >= len(nums)-1 {
			return true
		}

	}
	return false

}
