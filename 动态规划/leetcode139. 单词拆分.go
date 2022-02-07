package main

func wordBreak(s string, wordDict []string) bool {
	wordDictMap := make(map[string]bool)
	for _, w := range wordDict {
		wordDictMap[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]

}
