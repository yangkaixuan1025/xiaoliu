package main

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	permuteBackTracking(nums, []int{}, make([]bool, len(nums)))
	return res
}
func permuteBackTracking(nums, r []int, used []bool) {
	if len(r) == len(nums) {
		rr := make([]int, len(r))
		copy(rr, r)
		res = append(res, rr)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		r = append(r, nums[i])
		used[i] = true
		permuteBackTracking(nums, r, used)
		r = r[:len(r)-1]
		used[i] = false
	}

}
