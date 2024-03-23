package Solution

import "container/heap"

func Solution(heights []int, bricks int, ladders int) int {
	// 1  2  3  4  5  6
	// 4  0  1  1  1  9996  diff bsearchr+judge????
	// b=4, l=1
	diff := make([]int, len(heights))
	for i := 1; i < len(heights); i++ {
		if r := heights[i] - heights[i-1]; r > 0 {
			diff[i] = r
		}
	}
	l, r := 0, len(heights)
	for l < r {
		mid := (l + r) / 2
		tmp := make([]int, mid+1)
		for i := 0; i <= mid; i++ {
			tmp[i] = diff[i]
		}
		if isOk1642(tmp, bricks, ladders) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if l != 0 {
		l--
	}
	return l
}

type heap1642 []int

func (h *heap1642) Len() int {
	return len(*h)
}

func (h *heap1642) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}
func (h *heap1642) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *heap1642) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *heap1642) Pop() interface{} {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}
func isOk1642(diffs []int, b, l int) bool {
	h := heap1642(diffs)
	heap.Init(&h)
	for h.Len() > 0 {
		top := heap.Pop(&h).(int)
		if top <= 0 {
			continue
		}
		if top <= b {
			b -= top
			continue
		}
		if l > 0 {
			l--
			continue
		}
		return false
	}
	return true
}
