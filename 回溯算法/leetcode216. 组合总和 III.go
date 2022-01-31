package main

func combinationSum3(k int, n int) [][]int {
	res = make([][]int, 0)
	if k < 0 || n < 0 {
		return res
	}
	combinationSum3BackBackTrack(k, n, 1, 0, []int{})
	return res

}

func combinationSum3BackBackTrack(k, n, start, sum int, track []int) {
	if sum > n {
		return
	}
	if len(track) == k && n == sum {
		temp := make([]int, k)
		copy(temp, track)
		res = append(res, temp)
		return
	}
	for i := start; i <= 9-(k-len(track))+1; i++ {
		track = append(track, i)
		sum += i
		combinationSum3BackBackTrack(k, n, i+1, sum, track)
		sum -= i
		track = track[:len(track)-1]
	}

}
