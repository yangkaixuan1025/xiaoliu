package main

import "sort"

func permuteUnique(nums []int) [][]int {
	res = make([][]int, 0)
	sort.Ints(nums)

	permuteUniqueBackTracking(nums, []int{}, make([]bool, len(nums)))
	return res
}

func permuteUniqueBackTracking(nums, r []int, used []bool) {
	if len(r) == len(nums) {
		rr := make([]int, len(r))
		copy(rr, r)
		res = append(res, rr)
		return
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
			continue
		}
		if !used[i] {
			r = append(r, nums[i])
			used[i] = true
			permuteUniqueBackTracking(nums, r, used)
			r = r[:len(r)-1]
			used[i] = false
		}
	}
}
