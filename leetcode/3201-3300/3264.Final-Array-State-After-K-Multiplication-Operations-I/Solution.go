package Solution

import "container/heap"

type heap3264 [][2]int

func (h *heap3264) Len() int {
	return len(*h)
}

func (h *heap3264) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heap3264) Less(i, j int) bool {
	a, b := (*h)[i], (*h)[j]
	if a[1] == b[1] {
		return a[0] < b[0]
	}
	return a[1] < b[1]
}

func (h *heap3264) Push(x any) {
	*h = append(*h, x.([2]int))
}

func (h *heap3264) Pop() any {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

func Solution(nums []int, k int, multiplier int) []int {
	h := &heap3264{}
	for i := range nums {
		heap.Push(h, [2]int{i, nums[i]})
	}
	for ; k > 0; k-- {
		top := heap.Pop(h).([2]int)
		top[1] *= multiplier
		heap.Push(h, top)
	}
	for h.Len() > 0 {
		top := heap.Pop(h).([2]int)
		nums[top[0]] = top[1]
	}
	return nums
}
