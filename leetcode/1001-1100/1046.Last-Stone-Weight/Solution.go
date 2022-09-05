package Solution

import "container/heap"

type stoneList []int

func (s *stoneList) Len() int {
	return len(*s)
}

func (s *stoneList) Less(i, j int) bool {
	return (*s)[i] > (*s)[j]
}

func (s *stoneList) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *stoneList) Push(x interface{}) {
	*s = append(*s, x.(int))
}

func (s *stoneList) Pop() interface{} {
	old := *s
	l := len(old)
	x := old[l-1]
	*s = old[:l-1]
	return x
}

func Solution(stones []int) int {
	if len(stones) == 0 {
		return 0
	}
	h := stoneList(stones)
	h1 := &h
	heap.Init(h1)

	for h1.Len() > 1 {
		top := heap.Pop(h1).(int)
		second := heap.Pop(h1).(int)
		if top == second {
			continue
		}
		heap.Push(h1, top-second)
	}

	if h1.Len() == 0 {
		return 0
	}
	return heap.Pop(h1).(int)
}
