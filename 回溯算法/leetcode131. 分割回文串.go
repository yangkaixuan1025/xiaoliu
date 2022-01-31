package main

var resStr [][]string

func partition(s string) [][]string {
	resStr = make([][]string, 0)
	partitionBackTracking(s, 0, []string{})
	return resStr
}

func partitionBackTracking(s string, index int, r []string) {
	if len(s) == index {
		rr := make([]string, len(r))
		copy(rr, r)
		resStr = append(resStr, rr)
		return
	}
	for i := index; i < len(s); i++ {
		if isPartition(s, index, i) {
			r = append(r, s[index:i+1])
		} else {
			continue
		}
		partitionBackTracking(s, i+1, r)
		r = r[:len(r)-1]

	}

}

//判断是否为回文
func isPartition(s string, startIndex, end int) bool {
	left := startIndex
	right := end
	for left < right {
		if s[left] != s[right] {
			return false
		}
		//移动左右指针
		left++
		right--
	}
	return true
}
