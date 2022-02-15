package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type hp struct {
	sort.IntSlice
}

type MedianFinder struct {
	queMin, queMax hp
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

/** initialize your data structure here. */
func MConstructor() MedianFinder {
	return MedianFinder{}

}

func (this *MedianFinder) AddNum(num int) {
	minQ, maxQ := &this.queMin, &this.queMax
	if minQ.Len() == 0 || num <= -minQ.IntSlice[0] {
		heap.Push(minQ, -num)
		if maxQ.Len()+1 < minQ.Len() {
			heap.Push(maxQ, -heap.Pop(minQ).(int))
		}
	} else {
		heap.Push(maxQ, num)
		if maxQ.Len() > minQ.Len() {
			heap.Push(minQ, -heap.Pop(maxQ).(int))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	minQ, maxQ := this.queMin, this.queMax
	if minQ.Len() > maxQ.Len() {
		return float64(-minQ.IntSlice[0])
	}
	return float64(maxQ.IntSlice[0]-minQ.IntSlice[0]) / 2

}

func main() {
	h := &hp{
		sort.IntSlice{0, 0, 100, 1, 2, 4, 2, 2, 3, 1, -20000, 4},
	}
	heap.Init(h)
	fmt.Println(h.IntSlice)
}
