package main

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	vis := make([][]bool, h)
	for i := range vis {
		vis[i] = make([]bool, w)
	}
	for i, row := range board {
		for j := range row {
			if existDfs(board, word, i, j, 0, h, w, vis) {
				return true
			}

		}
	}
	return false
}

func existDfs(board [][]byte, word string, i, j, k, h, w int, vis [][]bool) bool {
	if board[i][j] != word[k] {
		return false
	}
	if len(word)-1 == k {
		return true
	}
	vis[i][j] = true
	for _, dir := range directions {
		if newI, newJ := i+dir.x, j+dir.y; newI >= 0 && newI < h && newJ >= 0 && newJ < w && !vis[newI][newJ] {
			if existDfs(board, word, newI, newJ, k+1, h, w, vis) {
				return true
			}
		}
	}
	vis[i][j] = false

	return false
}
