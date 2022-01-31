package main

var letterMap = map[rune]string{
	'0': "",
	'1': "",
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}
var result []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	result = make([]string, 0)
	digitsRune := []rune(digits)
	letterCombinationsBackTracking(0, len(digitsRune), []rune{}, digitsRune)
	return result

}

func letterCombinationsBackTracking(index, n int, s []rune, digits []rune) {
	if len(s) == n {
		var ss string
		ss = string(s)
		result = append(result, ss)
		return
	}
	letters := []rune(letterMap[digits[index]])
	for i := 0; i < len(letters); i++ {
		s = append(s, letters[i])
		letterCombinationsBackTracking(index+1, n, s, digits)
		s = s[:len(s)-1]
	}
}
