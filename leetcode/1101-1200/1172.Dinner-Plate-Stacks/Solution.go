package Solution

import "container/heap"

type heapItem1172 struct {
	stack []int

	index, sourceIndex int
}

type leftNotFullHeap struct {
	data     []*heapItem1172
	capacity int
}

func (h *leftNotFullHeap) Len() int {
	return len(h.data)
}

func (h *leftNotFullHeap) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	la, lb := len(a.stack), len(b.stack)
	if la == h.capacity {
		return false
	}
	if lb == h.capacity {
		return true
	}
	return a.sourceIndex < b.sourceIndex
}

func (h *leftNotFullHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
	h.data[i].index = i
	h.data[j].index = j
}

func (h *leftNotFullHeap) Push(x any) {
	item := x.(*heapItem1172)
	l := len(h.data)
	item.index = l
	h.data = append(h.data, item)
}

func (h *leftNotFullHeap) Pop() any {
	old := h.data
	l := len(old)
	x := old[l-1]
	h.data = old[:l-1]
	return x
}

type rightHeapItem1172 struct {
	sourceIndex, index int
}
type rightNotEmptyHeap struct {
	data []*rightHeapItem1172
	list *[]*heapItem1172
}

func (h *rightNotEmptyHeap) Len() int {
	return len(h.data)
}

func (h *rightNotEmptyHeap) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	ia, ib := (*h.list)[a.sourceIndex], (*h.list)[b.sourceIndex]
	li, lj := len(ia.stack), len(ib.stack)
	if li == 0 && lj == 0 || (li != 0 && lj != 0) {
		return a.sourceIndex > b.sourceIndex
	}
	return li != 0
}

func (h *rightNotEmptyHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
	h.data[i].index = i
	h.data[j].index = j
}

func (h *rightNotEmptyHeap) Push(x any) {
	item := x.(*rightHeapItem1172)
	item.index = len(h.data)
	h.data = append(h.data, item)
}

func (h *rightNotEmptyHeap) Pop() any {
	l := len(h.data)
	x := h.data[l-1]
	h.data = h.data[:l-1]
	return x
}

type DinnerPlates struct {
	leftNotFull   *leftNotFullHeap
	rightNotEmpty *rightNotEmptyHeap
	capacity      int
	leftIndies    *[]*heapItem1172

	rightIndies map[int]*rightHeapItem1172
}

func Constructor(capacity int) DinnerPlates {
	indies := make([]*heapItem1172, 0)
	return DinnerPlates{
		leftNotFull: &leftNotFullHeap{
			data:     make([]*heapItem1172, 0),
			capacity: capacity,
		},
		rightNotEmpty: &rightNotEmptyHeap{
			data: make([]*rightHeapItem1172, 0),
			list: &indies,
		},
		capacity:    capacity,
		leftIndies:  &indies,
		rightIndies: make(map[int]*rightHeapItem1172),
	}
}

func (this *DinnerPlates) Push(val int) {
	if len(this.leftNotFull.data) == 0 || len(this.leftNotFull.data[0].stack) == this.capacity {
		sourceIndex := len(this.leftNotFull.data)
		leftItem := &heapItem1172{
			stack:       []int{val},
			sourceIndex: sourceIndex,
		}
		rightItem := &rightHeapItem1172{
			sourceIndex: sourceIndex,
		}
		this.rightIndies[sourceIndex] = rightItem

		*this.leftIndies = append(*this.leftIndies, leftItem)
		heap.Push(this.leftNotFull, leftItem)
		heap.Push(this.rightNotEmpty, rightItem)
		return
	}
	zero := this.leftNotFull.data[0]
	zero.stack = append(zero.stack, val)
	heap.Fix(this.leftNotFull, 0)
	heap.Fix(this.rightNotEmpty, this.rightIndies[zero.sourceIndex].index)
}

func (this *DinnerPlates) Pop() int {
	top := this.rightNotEmpty.data[0]
	return this.PopAtStack(top.sourceIndex)
}

func (this *DinnerPlates) PopAtStack(index int) int {
	if index < 0 || index >= len(*this.leftIndies) {
		return -1
	}
	v := (*this.leftIndies)[index]
	l := len(v.stack)
	if l == 0 {
		return -1
	}
	x := v.stack[l-1]
	v.stack = v.stack[:l-1]
	heap.Fix(this.leftNotFull, v.index)
	heap.Fix(this.rightNotEmpty, this.rightIndies[v.sourceIndex].index)
	return x
}

type opt struct {
	name string
	i    int
}

func Solution(capacity int, opts []opt) []int {
	c := Constructor(capacity)
	var ret []int
	for _, op := range opts {
		if op.name == "push" {
			c.Push(op.i)
			continue
		}
		if op.name == "popAtStack" {
			ret = append(ret, c.PopAtStack(op.i))
			continue
		}
		if op.name == "pop" {
			ret = append(ret, c.Pop())
		}
	}
	return ret
}
