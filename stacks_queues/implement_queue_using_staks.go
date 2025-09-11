package stacks_queues

import (
	"slices"
)

type MyQueue struct {
	StackIn  []int
	StackOut []int
}

func Constructor() MyQueue {
	return MyQueue{
		StackIn:  []int{},
		StackOut: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.StackIn = append(this.StackIn, x)
}

func (this *MyQueue) Pop() int {
	if len(this.StackOut) == 0 {
		this.MoveInToOut()
	}
	top := this.StackOut[len(this.StackOut)-1]
	this.StackOut = this.StackOut[:len(this.StackOut)-1]
	return top
}

func (this *MyQueue) MoveInToOut() {
	this.StackOut = this.StackIn
	slices.Reverse(this.StackOut)
	this.StackIn = []int{}
}

func (this *MyQueue) Peek() int {
	if len(this.StackOut) == 0 {
		this.MoveInToOut()
	}
	return this.StackOut[len(this.StackOut)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.StackIn) == 0 && len(this.StackOut) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
