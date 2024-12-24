package Solution

import "container/heap"

type item2593 struct {
	v, i int
}

type heap2593 []item2593

func (h *heap2593) Len() int {
	return len(*h)
}

func (h *heap2593) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *heap2593) Less(i, j int) bool {
	a, b := (*h)[i], (*h)[j]
	if a.v == b.v {
		return a.i < b.i
	}
	return a.v < b.v
}

func (h *heap2593) Push(x any) {
	*h = append(*h, x.(item2593))
}
func (h *heap2593) Pop() any {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

func Solution(nums []int) int64 {
	l := len(nums)
	used := make([]bool, l)
	h := &heap2593{}
	for i := range l {
		heap.Push(h, item2593{v: nums[i], i: i})
	}
	ans := int64(0)
	for h.Len() > 0 {
		top := heap.Pop(h).(item2593)
		if used[top.i] {
			continue
		}
		ans += int64(top.v)
		if top.i-1 >= 0 {
			used[top.i-1] = true
		}
		if top.i+1 < l {
			used[top.i+1] = true
		}
	}
	return ans
}
