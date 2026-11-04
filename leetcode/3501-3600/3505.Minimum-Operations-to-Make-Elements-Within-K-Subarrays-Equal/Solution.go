package Solution

import "container/heap"

func Solution(nums []int, x int, k int) int64 {
	n := len(nums)
	left := &intHeap{isMax: true}
	right := &intHeap{}
	delayed := map[int]int{}
	leftSize, rightSize := 0, 0
	var sumL, sumR int64
	ops := make([]int64, n)

	prune := func(h *intHeap) {
		for h.Len() > 0 && delayed[h.data[0]] > 0 {
			delayed[h.data[0]]--
			heap.Pop(h)
		}
	}

	balance := func() {
		if leftSize > rightSize+1 {
			num := heap.Pop(left).(int)
			heap.Push(right, num)
			sumR += int64(num)
			sumL -= int64(num)
			rightSize++
			leftSize--
			prune(left)
		} else if rightSize > leftSize {
			num := heap.Pop(right).(int)
			heap.Push(left, num)
			sumL += int64(num)
			sumR -= int64(num)
			leftSize++
			rightSize--
			prune(right)
		}
	}

	for i, num := range nums {
		if left.Len() == 0 || num <= left.data[0] {
			heap.Push(left, num)
			leftSize++
			sumL += int64(num)
		} else {
			heap.Push(right, num)
			rightSize++
			sumR += int64(num)
		}
		balance()
		if i >= x {
			num := nums[i-x]
			delayed[num]++
			if num <= left.data[0] {
				leftSize--
				sumL -= int64(num)
				if num == left.data[0] {
					prune(left)
				}
			} else {
				rightSize--
				sumR -= int64(num)
				if num == right.data[0] {
					prune(right)
				}
			}
			balance()
		}
		if i >= x-1 {
			mid := int64(left.data[0])
			ops[i] = sumR - sumL + mid*int64(leftSize-rightSize)
		}
	}

	dp := make([][]int64, n)
	for i := range dp {
		dp[i] = make([]int64, k+1)
		for j := 1; j <= k; j++ {
			dp[i][j] = 1<<63 - 1
		}
	}
	for j := 1; j <= k; j++ {
		if j == 1 {
			dp[x-1][j] = ops[x-1]
		} else {
			dp[j*x-1][j] = dp[(j-1)*x-1][j-1] + ops[j*x-1]
		}
	}
	for j := 1; j <= k; j++ {
		for i := j * x; i < n; i++ {
			dp[i][j] = min(dp[i-x][j-1]+ops[i], dp[i-1][j])
		}
	}
	return dp[n-1][k]
}

type intHeap struct {
	data  []int
	isMax bool
}

func (h intHeap) Len() int { return len(h.data) }
func (h intHeap) Less(i, j int) bool {
	if h.isMax {
		return h.data[i] > h.data[j]
	}
	return h.data[i] < h.data[j]
}
func (h intHeap) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *intHeap) Push(v any)   { h.data = append(h.data, v.(int)) }
func (h *intHeap) Pop() any {
	n := len(h.data)
	v := h.data[n-1]
	h.data = h.data[:n-1]
	return v
}
