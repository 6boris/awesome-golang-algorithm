package Solution

import "container/heap"

type heap1792 [][]int

func (h *heap1792) Len() int {
	return len(*h)
}

func (h *heap1792) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heap1792) Less(i, j int) bool {
	a, b := (*h)[i], (*h)[j]
	// 判断增幅
	aa := (float64(a[0]+1) / float64(a[1]+1)) - (float64(a[0]) / float64(a[1]))
	bb := (float64(b[0]+1) / float64(b[1]+1)) - (float64(b[0]) / float64(b[1]))
	return aa > bb
}

func (h *heap1792) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *heap1792) Pop() any {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

func Solution(classes [][]int, extraStudents int) float64 {
	h := &heap1792{}
	var ans float64 = 0
	for _, c := range classes {
		if c[0] == c[1] {
			// 永远都是1提升不了
			ans++
			continue
		}
		heap.Push(h, c)
	}
	if h.Len() > 0 {
		for ; extraStudents > 0; extraStudents-- {
			top := heap.Pop(h).([]int)
			top[0]++
			top[1]++
			heap.Push(h, top)
		}
		for h.Len() > 0 {
			top := heap.Pop(h).([]int)
			ans += float64(top[0]) / float64(top[1])
		}
	}

	return ans / float64(len(classes))
}
