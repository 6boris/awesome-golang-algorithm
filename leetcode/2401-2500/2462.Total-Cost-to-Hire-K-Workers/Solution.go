package Solution

import (
	"container/heap"
)

type pair2462 struct {
	index, value int
}

type pairs2462 []pair2462

func (p *pairs2462) Len() int {
	return len(*p)
}

func (p *pairs2462) Less(i, j int) bool {
	a, b := (*p)[i], (*p)[j]
	if a.value == b.value {
		return a.index < b.index
	}
	return a.value < b.value
}

func (p *pairs2462) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *pairs2462) Push(x interface{}) {
	*p = append(*p, x.(pair2462))
}

func (p *pairs2462) Pop() interface{} {
	old := *p
	l := len(old)
	x := old[l-1]
	*p = old[:l-1]
	return x
}

func Solution(costs []int, k int, candidates int) int64 {
	left := len(costs)
	l, r := pairs2462{}, pairs2462{}
	li, ri := 0, left-1
	ans := int64(0)
	if candidates+candidates < left {
		for ; li < candidates; li, ri = li+1, ri-1 {
			heap.Push(&l, pair2462{li, costs[li]})
			heap.Push(&r, pair2462{ri, costs[ri]})
		}
		for k > 0 && ri >= li {
			a := heap.Pop(&l).(pair2462)
			b := heap.Pop(&r).(pair2462)
			if a.value == b.value {
				ans += int64(a.value)
				if a.index < b.index {
					heap.Push(&l, pair2462{li, costs[li]})
					heap.Push(&r, b)
					li++
				} else {
					heap.Push(&r, pair2462{ri, costs[ri]})
					heap.Push(&l, a)
					ri--
				}
			} else if a.value < b.value {
				ans += int64(a.value)
				heap.Push(&l, pair2462{li, costs[li]})
				heap.Push(&r, b)
				li++
			} else {
				ans += int64(b.value)
				heap.Push(&r, pair2462{ri, costs[ri]})
				heap.Push(&l, a)
				ri--
			}
			k--
		}
	} else {
		for ; li < left; li++ {
			heap.Push(&l, pair2462{li, costs[li]})
		}
	}
	if k > 0 {
		for len(r) > 0 {
			a := heap.Pop(&r)
			heap.Push(&l, a)
		}
		for k > 0 && len(l) > 0 {
			k--
			a := heap.Pop(&l).(pair2462)
			ans += int64(a.value)
		}
	}

	return ans
}
