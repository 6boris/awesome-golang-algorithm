package Solution

import "container/heap"

type heapNode struct {
	maxScore, idx int
}
type MyHeap []heapNode

func (h MyHeap) Len() int {
	return len(h)
}

func (h MyHeap) Less(i, j int) bool {
	return h[i].maxScore > h[j].maxScore
}

func (h MyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MyHeap) Push(val interface{}) {
	*h = append(*h, val.(heapNode))
}

func (h *MyHeap) Pop() interface{} {
	l := len(*h)
	ret := (*h)[l-1]
	*h = (*h)[:l-1]
	return ret
}

func (h *MyHeap) Top() interface{} {
	if len(*h) > 0 {
		top := (*h)[0]
		return top
	}
	return nil
}

func Solution(nums []int, k int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	dp := make([]int, length)
	dp[length-1] = nums[length-1]
	priorityQueue := &MyHeap{}
	heap.Init(priorityQueue)
	heap.Push(priorityQueue, heapNode{maxScore: nums[length-1], idx: length - 1})
	for idx := length - 2; idx >= 0; idx-- {
		for priorityQueue.Len() > 0 && priorityQueue.Top().(heapNode).idx > idx+k {
			heap.Pop(priorityQueue)
		}
		dp[idx] = priorityQueue.Top().(heapNode).maxScore + nums[idx]
		heap.Push(priorityQueue, heapNode{maxScore: dp[idx], idx: idx})
		/*
			这个是最开始的想法，超时了。。
			firstSet := true // default is zero ,but maybe set zero
			for step := 1; step <= k && idx+step < length; step++ {
				score := nums[idx] + dp[idx+step]
				if firstSet || dp[idx] < score {
					dp[idx] = score
					firstSet = false
					continue
				}
			}
		*/
	}
	return dp[0]
}
