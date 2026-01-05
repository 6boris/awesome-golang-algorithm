package Solution

import (
	"container/heap"
)

type Item2931 struct {
	i, j int
	val  int64
}

type Heap2931 struct {
	data []Item2931
}

func (h *Heap2931) Len() int {
	return len(h.data)
}

func (h *Heap2931) Less(i, j int) bool {
	return h.data[i].val < h.data[j].val
}

func (h *Heap2931) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap2931) Push(x any) {
	h.data = append(h.data, x.(Item2931))
}

func (h *Heap2931) Pop() any {
	l := len(h.data)
	x := h.data[l-1]
	h.data = h.data[:l-1]
	return x
}

func Solution(values [][]int) int64 {
	h := Heap2931{data: make([]Item2931, 0)}
	rows, cols := len(values), len(values[0])
	for r := 0; r < rows; r++ {
		heap.Push(&h, Item2931{i: r, j: cols - 1, val: int64(values[r][cols-1])})
	}
	cnt := (cols - 1) * rows
	day := int64(1)

	var (
		ret  int64
		next Item2931
	)

	for cnt > 0 {
		top := heap.Pop(&h).(Item2931)
		ret += top.val * day
		if top.j != 0 {
			next = Item2931{i: top.i, j: top.j - 1, val: int64(values[top.i][top.j-1])}
			heap.Push(&h, next)
			cnt--
		}
		day++
	}
	for h.Len() > 0 {
		top := heap.Pop(&h).(Item2931)
		ret += top.val * day
		day++
	}
	return ret
}
