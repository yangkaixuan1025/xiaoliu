package main

type MyStack struct {
	q1 []int
	q2 []int
}

//
//func Constructor() MyStack {
//	return MyStack{
//		q1: make([]int, 0),
//		q2: make([]int, 0),
//	}
//}

func (this *MyStack) Push(x int) {
	this.q1 = append(this.q1, x)
	return
}

func (this *MyStack) Pop() int {
	if len(this.q1) == 0 {
		return 0
	}
	val := 0
	for i, v := range this.q1 {
		if i == len(this.q1)-1 {
			val = v
		} else {
			this.q2 = append(this.q2, v)
		}
	}
	this.q1 = this.q2
	this.q2 = make([]int, 0)
	return val
}

func (this *MyStack) Top() int {
	if len(this.q1) == 0 {
		return 0
	}
	val := 0
	for i, v := range this.q1 {
		if i == len(this.q1)-1 {
			val = v
		}
	}
	return val
}

func (this *MyStack) Empty() bool {
	return len(this.q1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
