package Solution

import "container/heap"

type heap2558 []int

func (h *heap2558) Len() int {
	return len(*h)
}

func (h *heap2558) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heap2558) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *heap2558) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *heap2558) Pop() any {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

func Sqrt2558(x int) int {
	if x == 0 {
		return 0
	}

	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}

	return r
}

func Solution(gifts []int, k int) int64 {
	ans := int64(0)
	h := &heap2558{}
	for _, n := range gifts {
		heap.Push(h, n)
	}
	for ; k > 0; k-- {
		top := heap.Pop(h).(int)
		x := Sqrt2558(top)
		heap.Push(h, x)
	}
	for h.Len() > 0 {
		top := heap.Pop(h).(int)
		ans += int64(top)
	}
	return ans
}
