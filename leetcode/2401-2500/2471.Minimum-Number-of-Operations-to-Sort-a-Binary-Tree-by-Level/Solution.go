package Solution

import (
	"sort"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func minimumCostSort(nums []int) int {
	b := make([]int, len(nums))
	rightPos := make(map[int]int)
	copy(b, nums)
	sort.Ints(b)
	for i, n := range b {
		rightPos[n] = i
	}
	swapCount := 0
	idx := 0
	for idx < len(nums) {
		if nums[idx] == b[idx] {
			idx++
			continue
		}
		correctPos := rightPos[nums[idx]]
		nums[idx], nums[correctPos] = nums[correctPos], nums[idx]
		swapCount++

	}
	/*
		for idx := 0; idx < len(nums); idx++ {
			if nums[idx] == b[idx] {
				continue
			}
			correctPos := rightPos[nums[idx]]
			nums[idx], nums[correctPos] = nums[correctPos], nums[idx]
			swapCount++
		}
	*/
	return swapCount
}

func Solution(root *TreeNode) int {
	ans := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		next := make([]*TreeNode, 0)
		values := make([]int, 0)
		for _, item := range queue {
			if item.Left != nil {
				next = append(next, item.Left)
				values = append(values, item.Left.Val)
			}
			if item.Right != nil {
				next = append(next, item.Right)
				values = append(values, item.Right.Val)
			}
		}
		ans += minimumCostSort(values)
		queue = next
	}
	return ans
}
