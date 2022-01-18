package main

func minSubArrayLen(target int, nums []int) int {
	result := len(nums) + 1
	sum := 0
	j := 0
	for i, v := range nums {
		sum += v
		for sum >= target {
			subLength := i - j + 1
			if subLength < result {
				result = subLength
			}
			sum -= nums[j]
			j++
		}
	}
	if result == len(nums)+1 {
		return 0
	} else {
		return result
	}

}
