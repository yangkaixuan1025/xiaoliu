package main

func climbStairs(n int) int {
	r := 0
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	pre := 2
	prePre := 1
	for i := 3; i <= n; i++ {
		r = pre + prePre
		prePre = pre
		pre = r
	}

	return r
}
