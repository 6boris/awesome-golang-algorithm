package Solution

import (
	"container/list"
)

type MaxQueue struct {
	queue *list.List
	deque *list.List
}

func Constructor() MaxQueue {
	return MaxQueue{
		queue: list.New(),
		deque: list.New(),
	}
}

func (this *MaxQueue) Max_value() int {
	if this.queue.Len() > 0 {
		return this.deque.Front().Value.(int)
	}
	return -1
}

func (this *MaxQueue) Push_back(value int) {
	this.queue.PushBack(value)

	//	保持 queue 单调递减
	for this.deque.Len() > 0 && this.deque.Back().Value.(int) < value {
		this.deque.Remove(this.deque.Back())
	}
	this.deque.PushBack(value)
}

func (this *MaxQueue) Pop_front() int {
	head := &list.Element{Value: -1}
	if this.queue.Len() > 0 {
		head = this.queue.Front()
		this.queue.Remove(this.queue.Front())
	}
	//	移除 deque 的值
	if this.deque.Len() > 0 && this.deque.Front().Value == head.Value {
		this.deque.Remove(this.deque.Front())
	}
	return head.Value.(int)
}
