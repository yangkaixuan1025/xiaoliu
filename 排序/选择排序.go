package main

func SelectionSort(s []int) {
	n := len(s)
	for i := 0; i < n; i++ {
		k := i
		for j := i + 1; j < n; j++ {
			if s[j] < s[k] {
				k = j
			}
		}
		Swap(s, i, k)
	}
}
