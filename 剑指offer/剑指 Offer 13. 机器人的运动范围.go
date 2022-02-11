package main

func movingCount(m int, n int, k int) int {
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	return movingCountDfs(m, n, 0, 0, k, vis)
}

func movingCountDfs(m, n, i, j, k int, vis [][]bool) int {
	count := 0
	for movingCountCheck(m, n, i, j, k, vis) {
		vis[i][j] = true
		count = 1 + movingCountDfs(m, n, i-1, j, k, vis) + movingCountDfs(m, n, i+1, j, k, vis) + movingCountDfs(m, n, i, j-1, k, vis) + movingCountDfs(m, n, i, j+1, k, vis)
	}

	return count
}

func movingCountCheck(m, n, i, j, k int, vis [][]bool) bool {
	if i >= 0 && i < m &&
		j >= 0 && j < n &&
		getDigitSum(i)+getDigitSum(j) <= k &&
		!vis[i][j] {
		return true
	}
	return false

}

func getDigitSum(number int) int {
	sum := 0
	for number > 0 {
		sum += number % 10
		number /= 10
	}
	return sum
}
