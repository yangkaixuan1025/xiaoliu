package main

func lemonadeChange(bills []int) bool {
	//left表示还剩多少 下标0位5元的个数 ，下标1为10元的个数
	left := [2]int{0, 0}
	//第一个元素不为5，直接退出
	if bills[0] != 5 {
		return false
	}
	for i := 0; i < len(bills); i++ {
		//先统计5元和10元的个数
		if bills[i] == 5 {
			left[0] += 1
		}
		if bills[i] == 10 {
			left[1] += 1
		}
		//接着处理找零的
		tmp := bills[i] - 5
		if tmp == 5 {
			if left[0] > 0 {
				left[0] -= 1
			} else {
				return false
			}
		}
		if tmp == 15 {
			if left[1] > 0 && left[0] > 0 {
				left[0] -= 1
				left[1] -= 1
			} else if left[1] == 0 && left[0] > 2 {
				left[0] -= 3
			} else {
				return false
			}
		}
	}
	return true
}
