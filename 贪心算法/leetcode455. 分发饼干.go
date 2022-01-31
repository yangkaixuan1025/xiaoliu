package main

import "sort"

func findContentChildren(g []int, s []int) int {
	n := 0
	if len(s) == 0 {
		return n
	}
	sort.Ints(g)
	sort.Ints(s)
	for i := 0; i < len(s); i++ {
		if n < len(g) && g[n] <= s[i] {
			n++
		}

	}

	return n
}
