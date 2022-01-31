package main

func combinationSum(candidates []int, target int) [][]int {
	res = make([][]int, 0)
	if len(candidates) == 0 {
		return res
	}
	combinationSumBackTracking(candidates, []int{}, 0, target, 0)
	return res

}

func combinationSumBackTracking(candidates, r []int, index, target, sum int) {
	if sum > target {
		return
	}
	if sum == target {
		rr := make([]int, len(r))
		copy(rr, r)
		res = append(res, rr)
		return
	}
	for i := index; i < len(candidates); i++ {
		r = append(r, candidates[i])
		sum += candidates[i]
		combinationSumBackTracking(candidates, r, i, target, sum)
		sum -= candidates[i]
		r = r[:len(r)-1]
	}

}
