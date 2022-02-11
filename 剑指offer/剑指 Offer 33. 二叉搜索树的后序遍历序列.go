package main

func verifyPostorder(postorder []int) bool {
	return verifyPostorderDfs(postorder, 0, len(postorder)-1)
}

func verifyPostorderDfs(postorder []int, i, j int) bool {
	if i >= j {
		return true
	}
	p := i
	for postorder[p] < postorder[j] {
		p++
	}
	m := p
	for postorder[p] > postorder[j] {
		p++
	}
	return p == j && verifyPostorderDfs(postorder, i, m-1) && verifyPostorderDfs(postorder, m, j-1)
}
