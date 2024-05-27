package Solution

import "container/list"

type MyCircularDeque struct {
	k    int
	list *list.List
}

func Constructor641(k int) MyCircularDeque {
	return MyCircularDeque{
		k:    k,
		list: list.New(),
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.list.Len() == this.k {
		return false
	}
	this.list.PushFront(value)
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.list.Len() == this.k {
		return false
	}
	this.list.PushBack(value)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.list.Len() == 0 {
		return false
	}
	head := this.list.Front()
	this.list.Remove(head)
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.list.Len() == 0 {
		return false
	}
	head := this.list.Back()
	this.list.Remove(head)
	return true

}

func (this *MyCircularDeque) GetFront() int {
	if this.list.Len() == 0 {
		return -1
	}
	return this.list.Front().Value.(int)
}

func (this *MyCircularDeque) GetRear() int {
	if this.list.Len() == 0 {
		return -1
	}
	return this.list.Back().Value.(int)

}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.list.Len() == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.k == this.list.Len()
}

type op struct {
	n string
	v int
}

func Solution(n int, opts []op) []any {
	c := Constructor641(n)
	ans := make([]any, 0)
	for _, o := range opts {
		if o.n == "if" {
			ans = append(ans, c.InsertFront(o.v))
			continue
		}
		if o.n == "il" {
			ans = append(ans, c.InsertLast(o.v))
			continue
		}
		if o.n == "df" {
			ans = append(ans, c.DeleteFront())
			continue
		}
		if o.n == "dl" {
			ans = append(ans, c.DeleteLast())
			continue
		}
		if o.n == "gf" {
			ans = append(ans, c.GetFront())
			continue
		}
		if o.n == "gr" {
			ans = append(ans, c.GetRear())
			continue
		}
		if o.n == "ie" {
			ans = append(ans, c.IsEmpty())
			continue
		}
		ans = append(ans, c.IsFull())
	}
	return ans
}
