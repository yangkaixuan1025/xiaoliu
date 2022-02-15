package main

import "strconv"

func translateNum(num int) int {

	src := strconv.FormatInt(int64(num), 10)
	switch len(src) {
	case 1:
		return 1

	}
	dp := make([]int, len(src))
	dp[0] = 1
	dp[1] = 1
	o := string(src[0:2])
	if o >= "10" && o <= "25" {
		dp[1] += 1
	}
	for i := 2; i < len(src); i++ {
		dp[i] = dp[i-1]
		o := string(src[i-1 : i+1])
		if o >= "10" && o <= "25" {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(src)-1]
}
