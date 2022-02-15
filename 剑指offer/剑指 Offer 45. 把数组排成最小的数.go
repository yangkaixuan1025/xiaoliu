package main

import (
	"fmt"
	"sort"
)

func minNumber(nums []int) string {

	sort.SliceStable(nums, func(i, j int) bool {
		if nums[i] == 0 {
			return true
		}
		if nums[j] == 0 {
			return false
		}
		x, y := nums[i], nums[j]
		nx, ny := 1, 1
		for x/nx > 0 {
			nx *= 10
		}
		for y/ny > 0 {
			ny *= 10
		}
		fmt.Println(x, nx, y, ny)
		return x*ny+y < y*nx+x

	})

	fmt.Println(nums)
	r := ""
	for _, num := range nums {
		r += fmt.Sprintf("%d", num)
	}
	return r

}

func main() {
	minNumber([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
}
