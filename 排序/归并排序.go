package main

func MergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := (r + l) / 2
	MergeSort(arr, l, mid)
	MergeSort(arr, mid+1, r)
	merge(arr, l, mid, r)
}

func merge(arr []int, l, mid, r int) {
	temp := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		temp[i-l] = arr[i]
	}

	left := l
	right := mid + 1
	for i := l; i <= r; i++ {
		if left > mid {
			arr[i] = temp[right-l]
			right++
		} else if right > r {
			arr[i] = temp[left-l]
			left++
		} else if temp[left-l] > temp[right-l] {
			arr[i] = temp[right-l]
			right++
		} else {
			arr[i] = temp[left-l]
			left++

		}
	}

}
