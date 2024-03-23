package Solution

import (
	"container/heap"
)

type heap1090 struct {
	values []int
	data   []int
}

func (h *heap1090) Len() int {
	return len(h.data)
}

func (h *heap1090) Less(i, j int) bool {
	return h.values[h.data[i]] > h.values[h.data[j]]
}
func (h *heap1090) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
func (h *heap1090) Push(x interface{}) {
	h.data = append(h.data, x.(int))
}
func (h *heap1090) Pop() interface{} {
	old := h.data
	l := len(old)
	x := old[l-1]
	h.data = old[:l-1]
	return x
}

func Solution(values []int, labels []int, numWanted int, useLimit int) int {
	group := make(map[int]int)
	h := heap1090{data: make([]int, 0), values: values}
	for i := 0; i < len(values); i++ {
		heap.Push(&h, i)
	}
	ans := 0
	for h.Len() > 0 && numWanted > 0 {
		top := heap.Pop(&h).(int)
		if group[labels[top]] >= useLimit {
			continue
		}
		ans += values[top]
		group[labels[top]]++
		numWanted--
	}
	return ans
}
