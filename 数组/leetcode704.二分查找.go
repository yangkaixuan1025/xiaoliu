package main

import "fmt"

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 13))
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1 // 区间：[left, right]
	for left <= right {
		mid := (right-left)/2 + left // 防止溢出
		num := nums[mid]
		if num == target {
			return mid
		}
		if num > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1

}

func searchVersion2(nums []int, target int) int {
	high := len(nums) // 区间:[low, high)
	low := 0
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid //  区间：[low, mid)
		} else {
			low = mid + 1
		}
	}
	return -1
}
