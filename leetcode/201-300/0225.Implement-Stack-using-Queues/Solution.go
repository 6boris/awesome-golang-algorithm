package Solution

type MyStack struct {
	enqueue []int
	dequeue []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	var mystack MyStack
	return mystack
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.enqueue = append(this.enqueue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	var value int
	if len(this.dequeue) == 0 {
		this.dequeue = this.enqueue
		this.enqueue = []int{}

		// 交换队列中元素
		for i, j := 0, len(this.dequeue)-1; i < len(this.dequeue) && j > i; i, j = i+1, j-1 {
			this.dequeue[i], this.dequeue[j] = this.dequeue[j], this.dequeue[i]
		}
	}
	if len(this.dequeue) > 0 {
		value = this.dequeue[0]
		if len(this.dequeue) == 1 {
			this.dequeue = []int{}
		} else {
			this.dequeue = this.dequeue[1:]
		}

	}
	return value
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if len(this.dequeue) > 0 {
		return this.dequeue[0]
	} else if len(this.enqueue) > 0 {
		return this.enqueue[len(this.enqueue)-1]
	}

	return 0
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.enqueue) == 0 && len(this.dequeue) == 0 {
		return true
	}
	return false
}

type MyStack1 struct {
	q1 []int
}

func Constructor225() MyStack1 {
	return MyStack1{
		q1: make([]int, 0),
	}
}

func (this *MyStack1) Push(x int) {
	tmp := make([]int, len(this.q1))
	// 先将queue清空，按顺序放到另一个queue中
	copy(tmp, this.q1)
	this.q1 = []int{x}
	// 原有元素入队
	this.q1 = append(this.q1, tmp...)
}

func (this *MyStack1) Pop() int {
	x := this.q1[0]
	this.q1 = this.q1[1:]
	return x
}

func (this *MyStack1) Top() int {
	return this.q1[0]
}

func (this *MyStack1) Empty() bool {
	return len(this.q1) == 0
}
