package Solution

import (
	"container/heap"
	"sort"
)

type heap502Item struct {
	cost, profit int
}

type heap502ItemList []heap502Item

func (h *heap502ItemList) Len() int {
	return len(*h)
}

func (h *heap502ItemList) Less(i, j int) bool {
	a, b := (*h)[i], (*h)[j]
	if a.profit == b.profit {
		return a.cost < b.cost
	}
	return a.profit > b.profit
}

func (h *heap502ItemList) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heap502ItemList) Push(x interface{}) {
	*h = append(*h, x.(heap502Item))
}

func (h *heap502ItemList) Pop() interface{} {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

func Solution(k int, w int, profits []int, capital []int) int {
	x := make([]heap502Item, len(profits))
	for i := 0; i < len(profits); i++ {
		x[i] = heap502Item{cost: capital[i], profit: profits[i]}
	}

	sort.Slice(x, func(i, j int) bool {
		return x[i].cost < x[j].cost
	})
	h := heap502ItemList{}
	pick := 0
	for ; k > 0; k-- {
		for ; pick < len(x) && x[pick].cost <= w; pick++ {
			heap.Push(&h, x[pick])
		}
		if len(h) == 0 {
			break
		}
		x := heap.Pop(&h).(heap502Item)
		w += x.profit
	}
	return w
}
