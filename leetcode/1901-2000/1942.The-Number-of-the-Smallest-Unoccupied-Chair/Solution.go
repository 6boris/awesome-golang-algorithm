package Solution

import (
	"container/heap"
	"sort"
)

type heap1942 struct {
	chair int
	end   int
}
type heap1942list []heap1942

func (h *heap1942list) Len() int {
	return len(*h)
}
func (h *heap1942list) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heap1942list) Less(i, j int) bool {
	a := (*h)[i].end
	b := (*h)[j].end
	if a == b {
		return (*h)[i].chair < (*h)[j].chair
	}
	return a < b
}

func (h *heap1942list) Push(x any) {
	*h = append(*h, x.(heap1942))
}

func (h *heap1942list) Pop() any {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

type chairs []int

func (c *chairs) Len() int {
	return len(*c)
}
func (c *chairs) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}
func (c *chairs) Less(i, j int) bool {
	return (*c)[i] < (*c)[j]
}
func (c *chairs) Push(x any) {
	*c = append(*c, x.(int))
}

func (c *chairs) Pop() any {
	old := *c
	l := len(old)
	x := old[l-1]
	*c = old[:l-1]
	return x
}

func Solution(times [][]int, targetFriend int) int {
	l := len(times)
	cc := &chairs{}
	for i := range l {
		heap.Push(cc, i)
	}
	list := make([][3]int, l)
	for i := range times {
		list[i] = [3]int{i, times[i][0], times[i][1]}
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i][1] < list[j][1]
	})
	usedChairs := &heap1942list{}

	for i := range list {
		for len(*usedChairs) > 0 && (*usedChairs)[0].end <= list[i][1] {
			top := heap.Pop(usedChairs).(heap1942)
			heap.Push(cc, top.chair)
		}
		chair := heap.Pop(cc).(int)
		if list[i][0] == targetFriend {
			return chair
		}
		heap.Push(usedChairs, heap1942{chair: chair, end: list[i][2]})
	}
	return -1
}
