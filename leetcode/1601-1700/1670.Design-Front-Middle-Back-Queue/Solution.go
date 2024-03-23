package Solution

import "container/list"

type FrontMiddleBackQueue struct {
	list *list.List
}

func Constructor1670() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		list: list.New(),
	}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.list.PushFront(val)
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	l := this.list.Len()
	if l == 0 {
		this.PushFront(val)
		return
	}

	mid := l / 2
	e := this.list.Front()
	for i := 0; i < mid; i++ {
		e = e.Next()
	}
	this.list.InsertBefore(val, e)
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.list.PushBack(val)
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.list.Len() == 0 {
		return -1
	}
	e := this.list.Front()
	v := e.Value.(int)
	this.list.Remove(e)
	return v
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	l := this.list.Len()
	if l == 0 {
		return -1
	}
	if l&1 == 0 {
		l--
	}
	mid := l / 2
	e := this.list.Front()
	for i := 0; i < mid; i++ {
		e = e.Next()
	}
	v := e.Value.(int)
	this.list.Remove(e)
	return v
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.list.Len() == 0 {
		return -1
	}
	e := this.list.Back()
	v := e.Value.(int)
	this.list.Remove(e)
	return v
}

type opt struct {
	name string
	val  int
}

func Solution(opts []opt) []int {
	c := Constructor1670()
	ans := make([]int, 0)
	for _, op := range opts {
		if op.name == "pushBack" {
			c.PushBack(op.val)
		}
		if op.name == "pushFront" {
			c.PushFront(op.val)
		}
		if op.name == "pushMiddle" {
			c.PushMiddle(op.val)
		}
		if op.name == "popFront" {
			ans = append(ans, c.PopFront())
		}
		if op.name == "popMiddle" {
			ans = append(ans, c.PopMiddle())
		}
		if op.name == "popBack" {
			ans = append(ans, c.PopBack())
		}
	}
	return ans
}
