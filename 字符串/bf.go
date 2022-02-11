package main

func bfSearch(main, pattern string) int {
	if len(main) == 0 || len(pattern) == 0 || len(pattern) > len(main) {
		return -1
	}
	for i := 0; i <= len(main)-len(pattern); i++ {
		if main[i:i+len(pattern)] == pattern {
			return i
		}
	}

	return -1

}
