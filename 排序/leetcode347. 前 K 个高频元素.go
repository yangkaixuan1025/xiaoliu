package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	ans := []int{}
	map_num := map[int]int{}
	for _, item := range nums {
		map_num[item]++
	}
	for key, _ := range map_num {
		ans = append(ans, key)
	}
	//核心思想：排序
	//可以不用包函数，自己实现快排
	sort.Slice(ans, func(a, b int) bool {
		return map_num[ans[a]] > map_num[ans[b]]
	})

	return ans[:k]
}
