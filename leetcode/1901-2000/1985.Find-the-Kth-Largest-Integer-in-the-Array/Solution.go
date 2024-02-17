package Solution

import (
	"container/heap"
)

type hp1985 []string

func greater1985(a, b string) bool {
	la, lb := len(a), len(b)
	if la > lb {
		return true
	} else if la < lb {
		return false
	}
	for k := 0; k < la; k++ {
		if a[k] == b[k] {
			continue
		}
		return a[k] > b[k]
	}
	return true
}
func (h *hp1985) Len() int {
	return len(*h)
}

func (h *hp1985) Less(i, j int) bool {
	return !greater1985((*h)[i], (*h)[j])
}

func (h *hp1985) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *hp1985) Push(x interface{}) {
	*h = append(*h, x.(string))
}

func (h *hp1985) Pop() interface{} {
	old := *h
	l := len(*h)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

func Solution(nums []string, k int) string {
	h := hp1985{}
	for _, str := range nums {
		if h.Len() == k {
			if !greater1985(h[0], str) {
				heap.Pop(&h)
				heap.Push(&h, str)
			}
			continue
		}
		heap.Push(&h, str)
	}
	return h[0]
}
