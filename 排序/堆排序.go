package main

func heapSort(a []int, n int) {
	buildHeap(a, n)
	k := n
	for k >= 1 {
		swap(a, 1, k)
		heapifyUpToDown(a, 1, k-1)
		k--
	}

}
func buildHeap(a []int, n int) {
	for i := n / 2; i > 1; i-- {
		heapifyUpToDown(a, i, n)
	}
}

func heapifyUpToDown(a []int, top int, count int) {
	for i := top; i <= count/2; {
		maxIndex := i
		if a[i] < a[i*2] {
			maxIndex = i * 2
		}
		if i*2+1 <= count && a[maxIndex] < a[i*2+1] {
			maxIndex = i*2 + 1
		}
		if maxIndex == i {
			break
		}
		swap(a, i, maxIndex)
		i = maxIndex
	}
}

//swap two elements
func swap(a []int, i int, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}
