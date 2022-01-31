package main

var res [][]int

func combine(n int, k int) [][]int {
	res = make([][]int, 0)
	if k > n || n < 0 || k < 0 {
		return res
	}
	backtrack(n, k, 1, []int{})
	return res
}

func backtrack(n, k, startIndex int, track []int) {
	if len(track) == k {
		temp := make([]int, k)
		copy(temp, track)
		res = append(res, temp)
		return
	}
	if len(track)+n-startIndex+1 < k {
		return
	}

	for i := startIndex; i <= n; i++ {
		track = append(track, i)
		backtrack(n, k, i+1, track)
		track = track[:len(track)-1]
	}
}
