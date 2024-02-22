package Solution

import "container/heap"

type h786 [][]int

func (h *h786) Len() int {
	return len(*h)
}

func (h *h786) Less(i, j int) bool {
	// a/b, c/d
	a := (*h)[i]
	b := (*h)[j]
	return a[0]*b[1] > a[1]*b[0]
}

func (h *h786) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *h786) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *h786) Pop() any {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

func Solution(arr []int, k int) []int {
	h := h786{}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if h.Len() != k {
				heap.Push(&h, []int{arr[i], arr[j]})
				continue
			}
			if h[0][0]*arr[j] > h[0][1]*arr[i] {
				heap.Pop(&h)
				heap.Push(&h, []int{arr[i], arr[j]})
			}
		}
	}
	return h[0]
}
