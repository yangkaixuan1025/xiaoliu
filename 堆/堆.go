package main

import (
	"container/heap"
	"fmt"
)

type Heap struct {
	a     []int
	n     int
	count int
}

func NewHeap(capacity int) *Heap {
	heap := &Heap{
		a:     make([]int, capacity+1),
		n:     capacity,
		count: 0,
	}
	return heap
}

// down to up
func (heap *Heap) insert(data int) {
	if heap.count == heap.n {
		return
	}
	heap.count++
	heap.a[heap.count] = data
	i := heap.count
	parent := i / 2
	for parent > 0 && heap.a[parent] < heap.a[i] {
		swap(heap.a, parent, i)
		i = parent
		parent = i / 2
	}
}

// up to down
func (heap *Heap) removeMax() {
	if heap.count == 0 {
		return
	}
	swap(heap.a, 1, heap.count)
	heap.count--

	heapifyUpToDown(heap.a, heap.count)

}

func heapifyUpToDown(a []int, count int) {
	for i := 1; i <= count/2; {
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

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getLeastNumbers(arr []int, k int) []int {

	h := IntHeap(arr[:k])
	hp := &h
	heap.Init(hp)

	for i := k; i < len(arr); i++ {
		fmt.Println(hp)
		x := heap.Pop(hp).(int)
		fmt.Println(x)
		if x > arr[i] {
			heap.Push(hp, arr[i])
		} else {
			heap.Push(hp, x)
		}
	}
	fmt.Println(hp)
	return []int(*hp)
}

func main() {
	//getLeastNumbers([]int{0, 0, 100, 1, 2, 4, 2, 2, 3, 1, 4}, 8)
	h := &IntHeap{0, 0, 100, 1, 2, 4, 2, 2, 3, 1, 20000, 4}
	heap.Init(h)
	fmt.Println(h)
}
