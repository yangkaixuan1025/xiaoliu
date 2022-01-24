package main

func InsertionSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i; a[j] < a[j-1] && j > 0; j-- {
			Swap(a, j, j-1)
		}
	}
}

func Swap(slice []int, i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
