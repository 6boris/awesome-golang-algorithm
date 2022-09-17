package Solution

import "container/heap"

type arr1502 []int

func (a *arr1502) Len() int {
	return len(*a)
}

func (a *arr1502) Less(i, j int) bool {
	return (*a)[i] < (*a)[j]
}

func (a *arr1502) Swap(i, j int) {
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
}

func (a *arr1502) Push(x interface{}) {
	*a = append(*a, x.(int))
}

func (a *arr1502) Pop() interface{} {
	old := *a
	l := len(old)
	x := old[l-1]
	*a = old[:l-1]
	return x
}

func Solution(arr []int) bool {
	h := arr1502(arr)
	heap.Init(&h)

	first := heap.Pop(&h).(int)
	diff := -1

	for (&h).Len() > 0 {
		second := heap.Pop(&h).(int)
		_diff := second - first
		if diff != -1 && diff != _diff {
			return false
		}
		diff = _diff
		first = second
	}
	return true
}
