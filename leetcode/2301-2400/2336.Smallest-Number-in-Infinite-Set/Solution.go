package Solution

import "container/heap"

type item2336 []int

func (r *item2336) Len() int {
	return len(*r)
}
func (r *item2336) Less(i, j int) bool {
	return (*r)[i] < (*r)[j]
}
func (r *item2336) Swap(i, j int) {
	(*r)[i], (*r)[j] = (*r)[j], (*r)[i]
}

func (r *item2336) Push(x interface{}) {
	*r = append(*r, x.(int))
}
func (r *item2336) Pop() interface{} {
	old := *r
	l := len(old)
	x := old[l-1]
	*r = old[:l-1]
	return x
}

type SmallestInfiniteSet struct {
	index int
	store item2336
}

func Constructor2336() SmallestInfiniteSet {
	return SmallestInfiniteSet{index: 1, store: item2336{}}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if len(this.store) != 0 {
		x := heap.Pop(&this.store).(int)
		return x
	}
	x := this.index
	this.index++
	return x
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.index {
		return
	}
	for _, n := range this.store {
		if n == num {
			return
		}
	}
	heap.Push(&this.store, num)
}

type Op struct {
	name string
	val  int
}

func Solution(ops []Op) []int {
	ans := make([]int, 0)
	c := Constructor2336()
	for _, o := range ops {
		if o.name == "addBack" {
			c.AddBack(o.val)
			continue
		}
		ans = append(ans, c.PopSmallest())
	}
	return ans
}
