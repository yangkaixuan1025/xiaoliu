package main

func removeElement(nums []int, val int) int {
	var slow int

	for i, v := range nums {
		if v != val {
			nums[i], nums[slow] = nums[slow], nums[i]
			slow++
		}
	}
	return slow
}

func removeElementVersion2(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}
