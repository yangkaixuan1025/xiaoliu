package main

func ShellSort(a []int) {
	l := len(a)
	for gap := l / 2; gap > 0; gap /= 2 {
		for i := gap; i < l; i++ {
			for j := i - gap; j > 0 && a[j] < a[j-gap]; j -= gap {
				Swap(a, j, j-gap)
			}
		}
	}
}
