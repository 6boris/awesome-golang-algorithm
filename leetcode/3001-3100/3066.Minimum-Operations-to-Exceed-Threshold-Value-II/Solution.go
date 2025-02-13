package Solution

import "container/heap"

type heap3066 []int

func (h *heap3066) Len() int {
	return len(*h)
}

func (h *heap3066) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *heap3066) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heap3066) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *heap3066) Pop() any {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

func Solution(nums []int, k int) int {
	steps := 0
	h := heap3066(nums)
	heap.Init(&h)
	var a, b int
	for h.Len() > 0 {
		top := h[0]
		if top >= k {
			break
		}
		a = heap.Pop(&h).(int)
		b = heap.Pop(&h).(int)
		heap.Push(&h, min(a, b)*2+max(a, b))
		steps++
	}
	return steps
}
