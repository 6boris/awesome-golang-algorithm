package Solution

import (
	"container/heap"
)

const mod1508 = 1000000007

type ints []int

func (i *ints) Len() int {
	return len(*i)
}

func (ii *ints) Swap(i, j int) {
	(*ii)[i], (*ii)[j] = (*ii)[j], (*ii)[i]
}

func (ii *ints) Less(i, j int) bool {
	return (*ii)[i] < (*ii)[j]
}

func (i *ints) Push(x interface{}) {
	*i = append(*i, x.(int))
}

func (i *ints) Pop() interface{} {
	old := *i
	l := len(old)
	x := old[l-1]
	*i = old[:l-1]
	return x
}

func Solution(nums []int, n int, left int, right int) int {
	h := ints{nums[0]}
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
		heap.Push(&h, nums[i])
		for pre := i; pre > 0; pre-- {
			heap.Push(&h, nums[i]-nums[pre-1])
		}
	}
	sum, index := 0, 0
	for len(h) > 0 {
		index++
		x := heap.Pop(&h).(int)
		if index < left {
			continue
		}
		if index > right {
			break
		}
		sum = (sum + x) % mod1508
	}
	return sum
}
