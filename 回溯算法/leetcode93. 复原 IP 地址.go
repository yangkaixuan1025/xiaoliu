package main

import (
	"fmt"
	"strconv"
)

var resS []string

func restoreIpAddresses(s string) []string {
	resS = make([]string, 0)
	restoreIpAddressesBackTracking(s, []string{}, 0)
	return resS
}

func restoreIpAddressesBackTracking(s string, r []string, index int) {
	if len(r) == 4 && index == len(s) {
		resS = append(resS, fmt.Sprintf("%s.%s.%s.%s", r[0], r[1], r[2], r[3]))
		return
	}
	for i := index; i < len(s); i++ {
		if isNormalIp(s, index, i) {
			r = append(r, s[index:i+1])
			restoreIpAddressesBackTracking(s, r, i+1)
		} else {
			continue
		}
		r = r[:len(r)-1]

	}

}

func isNormalIp(s string, startIndex, end int) bool {
	checkInt, _ := strconv.Atoi(s[startIndex : end+1])
	if end-startIndex+1 > 1 && s[startIndex] == '0' { //对于前导 0的IP（特别注意s[startIndex]=='0'的判断，不应该写成s[startIndex]==0，因为s截取出来不是数字）
		return false
	}
	if checkInt > 255 {
		return false
	}
	return true
}
