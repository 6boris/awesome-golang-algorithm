package Solution

import "container/heap"

func Solution(nums []int) int64 {
	n3 := len(nums)
	n := n3 / 3
	part1 := make([]int64, n+1)
	var sum int64 = 0
	ql := &MaxHeap{}
	heap.Init(ql)
	for i := 0; i < n; i++ {
		sum += int64(nums[i])
		heap.Push(ql, nums[i])
	}
	part1[0] = sum
	for i := n; i < n*2; i++ {
		sum += int64(nums[i])
		heap.Push(ql, nums[i])
		sum -= int64(heap.Pop(ql).(int))
		part1[i-(n-1)] = sum
	}

	var part2 int64 = 0
	qr := &IntMinHeap{}
	heap.Init(qr)
	for i := n*3 - 1; i >= n*2; i-- {
		part2 += int64(nums[i])
		heap.Push(qr, nums[i])
	}
	ans := part1[n] - part2
	for i := n*2 - 1; i >= n; i-- {
		part2 += int64(nums[i])
		heap.Push(qr, nums[i])
		part2 -= int64(heap.Pop(qr).(int))
		if part1[i-n]-part2 < ans {
			ans = part1[i-n] - part2
		}
	}
	return ans
}

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type IntMinHeap []int

func (h IntMinHeap) Len() int            { return len(h) }
func (h IntMinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntMinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
