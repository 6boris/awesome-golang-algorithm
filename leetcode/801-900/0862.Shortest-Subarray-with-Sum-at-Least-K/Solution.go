package Solution

import (
	"container/heap"
	"container/list"
	"math"
)

func Solution(nums []int, k int) int {
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	minLength := math.MaxInt32
	deque := list.New()

	for i := 0; i <= n; i++ {
		for deque.Len() > 0 && prefixSum[i]-prefixSum[deque.Front().Value.(int)] >= k {
			minLength = min(minLength, i-deque.Remove(deque.Front()).(int))
		}

		for deque.Len() > 0 && prefixSum[i] <= prefixSum[deque.Back().Value.(int)] {
			deque.Remove(deque.Back())
		}

		deque.PushBack(i)
	}

	if minLength == math.MaxInt32 {
		return -1
	}
	return minLength
}

type heap862 struct {
	v, i int
}
type heap862list []heap862

func (h *heap862list) Len() int {
	return len(*h)
}
func (h *heap862list) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heap862list) Less(i, j int) bool {
	return (*h)[i].v < (*h)[j].v
}
func (h *heap862list) Push(x any) {
	*h = append(*h, x.(heap862))
}
func (h *heap862list) Pop() any {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

func Solution1(nums []int, k int) int {
	h := heap862list{{0, -1}}
	cur := 0
	l := len(nums)
	ans := -1
	i := 0
	for ; i < l; i++ {
		cur += nums[i]
		for len(h) > 0 && cur-h[0].v >= k {
			top := heap.Pop(&h).(heap862)
			if ans == -1 || ans > i-top.i {
				ans = i - top.i
			}
		}
		heap.Push(&h, heap862{cur, i})
	}
	return ans
}
