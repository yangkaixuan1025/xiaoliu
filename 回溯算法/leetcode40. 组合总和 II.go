package main

// https://leetcode-cn.com/problems/combination-sum-ii/
import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	res = make([][]int, 0)
	sort.Ints(candidates)
	combinationSum2BackTracking(candidates, []int{}, target, 0, 0, make([]bool, len(candidates)))
	return res
}

func combinationSum2BackTracking(candidates, r []int, target, sum, index int, used []bool) {
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
		if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
			continue
		}
		sum += candidates[i]
		r = append(r, candidates[i])
		used[i] = true
		combinationSum2BackTracking(candidates, r, target, sum, i+1, used)
		sum -= candidates[i]
		r = r[:len(r)-1]
		used[i] = false
	}

}
