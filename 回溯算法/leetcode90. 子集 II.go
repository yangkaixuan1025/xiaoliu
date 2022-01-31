package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res = make([][]int, 0)
	sort.Ints(nums)
	subsetsWithDupBackTracking(nums, []int{}, 0)
	return res
}

func subsetsWithDupBackTracking(nums, r []int, index int) {
	rr := make([]int, len(r))
	copy(rr, r)
	res = append(res, rr)
	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		r = append(r, nums[i])
		subsetsWithDupBackTracking(nums, r, i+1)
		r = r[:len(r)-1]
	}

}
