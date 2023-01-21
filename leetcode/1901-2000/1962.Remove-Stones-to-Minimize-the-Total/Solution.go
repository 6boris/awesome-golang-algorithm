package Solution

import (
	"container/heap"
)

type Piles []int

func (p *Piles) Len() int {
	return len(*p)
}

func (p *Piles) Less(i, j int) bool {
	a := (*p)[i]
	b := (*p)[j]
	if a&1 == 1 {
		a++
	}
	if b&1 == 1 {
		b++
	}

	return (*p)[i]-a/2 > (*p)[j]-b/2
}

func (p *Piles) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *Piles) Push(x interface{}) {
	*p = append(*p, x.(int))
}

func (p *Piles) Pop() interface{} {
	old := *p
	l := len(old)
	x := old[l-1]
	*p = old[:l-1]
	return x
}

func Solution(piles []int, k int) int {
	sum := 0
	for _, p := range piles {
		sum += p
	}
	h := Piles(piles)
	heap.Init(&h)
	for ; k > 0 && h.Len() > 0; k-- {
		x := heap.Pop(&h).(int)
		y := x
		if x&1 == 1 {
			x--
		}
		sum -= x / 2
		heap.Push(&h, y-x/2)
	}

	return sum
}
