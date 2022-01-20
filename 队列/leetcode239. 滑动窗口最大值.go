package main

func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0)
	curr := make([]int, 0)
	currMax := 0
	for i, v := range nums {
		if (i+1)%k == 0 {
			curr = curr[1:]

		} else {
			if currMax <= v {
				currMax = v
			}
			curr = append(curr, v)
		}
	}
	return result
}
