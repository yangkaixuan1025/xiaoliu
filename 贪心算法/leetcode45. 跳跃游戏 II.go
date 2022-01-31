package main

func jump(nums []int) int {
	var res int

	cur := 0
	next := 0

	for i := 0; i < len(nums); i++ {
		if nums[i]+i > next {
			next = nums[i] + i
		}
		if i == cur {
			if cur != len(nums)-1 {
				res++
				cur = next
				if next >= len(nums)-1 {
					break
				}

			} else {
				break
			}

		}
	}

	return res

}
