package main

func findSubsequences(nums []int) [][]int {
	res = make([][]int, 0)
	findSubsequencesBackTracking(nums, []int{}, 0)
	return res
}

func findSubsequencesBackTracking(nums, r []int, index int) {
	if len(r) > 1 {
		rr := make([]int, len(r))
		copy(rr, r)
		res = append(res, rr)
	}
	usedMap := make(map[int]bool)
	for i := index; i < len(nums); i++ {
		if (len(r) > 0 && r[len(r)-1] > nums[i]) || usedMap[nums[i]] {
			continue
		}
		r = append(r, nums[i])
		usedMap[nums[i]] = true
		findSubsequencesBackTracking(nums, r, i+1)
		r = r[:len(r)-1]
	}

}
