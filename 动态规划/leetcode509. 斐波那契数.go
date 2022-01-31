package main

func fib(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	r := 0
	prePre := 0
	pre := 1
	for i := 2; i <= n; i++ {
		r = pre + prePre
		prePre = pre
		pre = r
	}
	return r

}
