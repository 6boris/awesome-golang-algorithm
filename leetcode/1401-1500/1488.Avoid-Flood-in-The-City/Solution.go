package Solution

import "container/heap"

// 记录未来即将出现的lake
type lakeState struct {
	lake    int
	nextDay int
}

type lakeHeap struct {
	data []*lakeState
}

func (l *lakeHeap) Len() int {
	return len(l.data)
}

func (l *lakeHeap) Swap(i, j int) {
	l.data[i], l.data[j] = l.data[j], l.data[i]
}

func (l *lakeHeap) Less(i, j int) bool {
	a, b := l.data[i], l.data[j]
	return a.nextDay < b.nextDay
}

func (l *lakeHeap) Push(x any) {
	l.data = append(l.data, x.(*lakeState))
}

func (l *lakeHeap) Pop() any {
	ll := len(l.data)
	x := l.data[ll-1]
	l.data = l.data[:ll-1]
	return x
}

func Solution(rains []int) []int {
	l := len(rains)
	days := make(map[int][]int)
	// 记录走到那一天了
	indies := make(map[int]int)
	for index, rain := range rains {
		if rain == 0 {
			continue
		}
		days[rain] = append(days[rain], index)
		indies[rain] = 0
	}
	full := make(map[int]bool)
	h := &lakeHeap{data: make([]*lakeState, 0)}
	ret := make([]int, l)
	for index, lake := range rains {
		if lake > 0 {
			if full[lake] {
				// 满
				return []int{}
			}
			full[lake] = true
			ret[index] = -1
			if nextIndex := indies[lake] + 1; nextIndex < len(days[lake]) {
				indies[lake]++
				heap.Push(h, &lakeState{lake: lake, nextDay: days[lake][nextIndex]})
			}
			continue
		}
		ret[index] = 1
		if h.Len() > 0 {
			top := heap.Pop(h).(*lakeState)
			ret[index] = top.lake
			delete(full, top.lake)
		}
	}
	// 1, 2, 0, 2, 3, 0, 1
	return ret
}
