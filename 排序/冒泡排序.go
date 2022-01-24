package main

import "fmt"

func main() {
	values := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	fmt.Println(values)
	BubbleSort(values)
}

func BubbleSort(values []int) {
	n := len(values)
	for i := 0; i < n; i++ {

		for j := 1; j < n-i; j++ {
			if values[j-1] > values[j] {
				values[j-1], values[j] = values[j], values[j-1]
			}

		}
	}
	fmt.Println(values)
}

func BubbleSortBest(values []int) {
	flag := len(values)
	k := flag
	for flag > 0 {
		k = flag
		flag = 0
		for i := 1; i < k; i++ {
			if values[i-1] > values[i] {
				values[i-1], values[i] = values[i], values[i-1]
				flag = i
			}
		}
	}
	fmt.Println(values)
}
