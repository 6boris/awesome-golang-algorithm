package Solution

import (
	"container/heap"
)

type heap1738 []int

func (h *heap1738) Len() int {
	return len(*h)
}

func (h *heap1738) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heap1738) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *heap1738) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *heap1738) Pop() interface{} {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

func Solution(matrix [][]int, k int) int {
	rows := len(matrix)
	cols := len(matrix[0])
	cache := make([][]int, rows+1)
	for i := 0; i <= rows; i++ {
		cache[i] = make([]int, cols+1)
	}
	h := heap1738{}
	for r := 1; r <= rows; r++ {
		for c := 1; c <= cols; c++ {
			cache[r][c] = matrix[r-1][c-1] ^ cache[r-1][c] ^ cache[r][c-1] ^ cache[r-1][c-1]
			if h.Len() != k {
				heap.Push(&h, cache[r][c])
				continue
			}
			top := heap.Pop(&h).(int)
			if top < cache[r][c] {
				top = cache[r][c]
			}
			heap.Push(&h, top)
		}
	}
	return h[0]
}
