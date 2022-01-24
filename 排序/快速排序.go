package main

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(arr, start, end)
	quickSort(arr, start, i-1)
	quickSort(arr, i+1, end)
}

func partition(arr []int, start, end int) int {
	pivot := arr[end]

	var i = start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
