package Solution

import "container/heap"

type i64score []int64

func (i *i64score) Len() int {
	return len(*i)
}

func (ii *i64score) Less(i, j int) bool {
	return (*ii)[i] > (*ii)[j]
}

func (ii *i64score) Swap(i, j int) {
	(*ii)[i], (*ii)[j] = (*ii)[j], (*ii)[i]
}

func (i *i64score) Push(x interface{}) {
	*i = append(*i, x.(int64))
}

func (i *i64score) Pop() interface{} {
	old := *i
	l := len(old)
	x := old[l-1]
	*i = old[:l-1]
	return x
}

func Solution(nums []int, k int) int64 {
	pq := i64score{}
	for _, n := range nums {
		heap.Push(&pq, int64(n))
	}
	ans := int64(0)
	for k > 0 && len(pq) > 0 {
		top := heap.Pop(&pq).(int64)
		ans += top
		a := top / 3
		if top%3 != 0 {
			a++
		}
		heap.Push(&pq, a)
		k--
	}
	return ans
}
