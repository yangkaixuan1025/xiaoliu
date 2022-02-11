package main

import (
	"fmt"
)

func main() {
	q := Constructor()
	q.Push(1)
	q.Push(1)
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())

}

type MyQueue struct {
	stack []int
	back  []int
}

func Constructor() MyQueue {
	return MyQueue{
		stack: make([]int, 0),
		back:  make([]int, 0),
	}

}

func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.back) == 0 && len(this.stack) == 0 {
		return 0
	}

	if len(this.back) == 0 {
		for i := len(this.stack) - 1; i >= 0; i-- {
			this.back = append(this.back, this.stack[i])
		}
		this.stack = make([]int, 0)
	}
	val := this.back[len(this.back)-1]
	this.back = this.back[:len(this.back)-1]
	return val

}

func (this *MyQueue) Peek() int {
	if len(this.back) == 0 && len(this.stack) == 0 {
		return 0
	}

	if len(this.back) == 0 {
		for i := len(this.stack) - 1; i >= 0; i-- {
			this.back = append(this.back, this.stack[i])
		}
		this.stack = make([]int, 0)
	}
	val := this.back[len(this.back)-1]
	return val
}

func (this *MyQueue) Empty() bool {
	return len(this.back) == 0 && len(this.stack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
