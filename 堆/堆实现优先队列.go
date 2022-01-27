package main

type Node struct {
	value    int
	priority int
}

type PQueue struct {
	heap []Node

	capacity int
	used     int
}

func NewPriorityQueue(capacity int) PQueue {
	return PQueue{
		heap:     make([]Node, capacity+1, capacity+1),
		capacity: capacity,
		used:     0,
	}
}

func (q *PQueue) Push(node Node) {
	if q.used > q.capacity {
		return
	}
	q.used++
	q.heap[q.used] = node
}

func (q *PQueue) Pop() Node {
	if q.used == 0 {
		return Node{
			value:    -1,
			priority: -1,
		}
	}
	// 先堆化, 再取堆顶元素
	adjustHeap(q.heap, 1, q.used)
	node := q.heap[1]
	q.heap[1] = q.heap[q.used]
	q.used--
	return node
}

func (q *PQueue) Top() Node {
	if q.used == 0 {
		return Node{
			value:    -1,
			priority: -1,
		}
	}

	adjustHeap(q.heap, 1, q.used)
	return q.heap[1]
}

func adjustHeap(src []Node, start, end int) {
	if start >= end {
		return
	}
	// 只需要保证优先级最高的节点在 src[i] 的位置上即可
	for i := end / 2; i >= start; i-- {
		high := i
		if src[high].priority < src[2*i].priority {
			high = 2 * i
		}
		if src[high].priority < src[2*i].priority {
			high = 2*i + 1
		}
		if high == i {
			continue
		}
		src[high], src[i] = src[i], src[high]
	}
}
