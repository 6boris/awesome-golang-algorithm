package Solution

import (
	"container/heap"
)

type h274 []int

func (h *h274) Len() int {
	return len(*h)
}

func (h *h274) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *h274) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *h274) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *h274) Pop() any {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

func Solution(citations []int) int {
	h := h274{}
	for _, n := range citations {
		heap.Push(&h, n)
	}
	cur := 0
	for h.Len() > 0 {
		top := heap.Pop(&h).(int)
		cur++
		if top < cur {
			cur--
			break
		}
	}
	return cur
}
