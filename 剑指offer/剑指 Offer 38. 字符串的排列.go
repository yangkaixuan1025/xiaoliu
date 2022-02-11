package main

import "sort"

var resStr []string

func permutation(s string) []string {
	resStr = make([]string, 0)
	ss := []rune(s)
	sort.Slice(ss, func(a, b int) bool {
		return ss[a] < ss[b]
	})
	s = string(ss)
	permutationDfs(s, "", make([]bool, len(s)))
	return resStr

}

func permutationDfs(s string, r string, used []bool) {
	if len(s) == len(r) {
		resStr = append(resStr, ""+r)
		return
	}
	for i, v := range s {
		if i > 0 && s[i] == s[i-1] && used[i-1] == false {
			continue
		}
		if used[i] == false {
			r += string(v)
			used[i] = true
			permutationDfs(s, r, used)
			r = r[:len(r)-1]
			used[i] = false
		}

	}
}
