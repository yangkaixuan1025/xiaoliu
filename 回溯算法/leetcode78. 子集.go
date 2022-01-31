package main

func subsets(nums []int) [][]int {
	res = make([][]int, 0)
	subsetsBackTracking(nums, []int{}, 0)
	return res
}

func subsetsBackTracking(nums, r []int, index int) {
	rr := make([]int, len(r))
	copy(rr, r)
	res = append(res, rr)
	for i := index; i < len(nums); i++ {
		r = append(r, nums[i])
		subsetsBackTracking(nums, r, i+1)
		r = r[:len(r)-1]
	}
}
