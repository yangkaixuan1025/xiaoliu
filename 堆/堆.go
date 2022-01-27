package main

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
