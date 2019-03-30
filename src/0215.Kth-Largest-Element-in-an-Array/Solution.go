package Solution

import (
	"container/heap"
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	positive := make([]int, 10001)
	negative := make([]int, 10001)
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			positive[nums[i]]++
		} else {
			negative[-nums[i]]++
		}
	}
	flag := k
	for i := 10000; i > -1; i-- {
		flag -= positive[i]
		if flag <= 0 {
			return i
		}
	}
	for i := 0; i < 10001; i++ {
		flag -= negative[i]
		if flag <= 0 {
			return -i
		}
	}
	return 0
}

func findKthLargest2(nums []int, k int) int {
	h := make(IntMinHeap, 0, len(nums))
	heap.Init(&h)

	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		heap.Init(&h)
	}
	for len(h) > 0 {
		fmt.Printf("%d ", h.Pop())
	}
	return 0
}
