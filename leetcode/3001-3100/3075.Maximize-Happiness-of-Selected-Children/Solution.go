package Solution

import "container/heap"

type list3075 []int

func (l *list3075) Len() int {
	return len(*l)
}

func (l *list3075) Less(i, j int) bool {
	return (*l)[i] > (*l)[j]
}
func (l *list3075) Swap(i, j int) {
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
}
func (l *list3075) Push(x interface{}) {
	*l = append(*l, x.(int))
}

func (l *list3075) Pop() interface{} {
	old := *l
	ll := len(old)
	x := old[ll-1]
	*l = old[:ll-1]
	return x
}
func Solution(happiness []int, k int) int64 {
	l := list3075(happiness)
	heap.Init(&l)
	remove := 0
	ans := int64(0)
	for k > 0 && len(l) > 0 {
		top := heap.Pop(&l).(int)
		diff := top - remove
		if diff <= 0 {
			break
		}
		ans += int64(diff)
		remove++
		k--

	}
	return ans
}
