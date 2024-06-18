package Solution

import (
	"math"
	"sort"
)

type SegmentTreeNode struct {
	Left, Right, Max      int
	LeftChild, RightChild *SegmentTreeNode
}

func buildSegmentTree(nums []workProfit, left, right int) *SegmentTreeNode {
	if left == right {
		return &SegmentTreeNode{Left: left, Right: right, Max: nums[left].p}
	}

	mid := (left + right) / 2
	leftNode := buildSegmentTree(nums, left, mid)
	rightNode := buildSegmentTree(nums, mid+1, right)
	max := max(leftNode.Max, rightNode.Max)

	return &SegmentTreeNode{Left: left, Right: right, Max: max, LeftChild: leftNode, RightChild: rightNode}
}

func queryMax(root *SegmentTreeNode, left, right int) int {
	if root.Left >= left && root.Right <= right {
		return root.Max
	}

	if root.Right < left || root.Left > right {
		return math.MinInt32
	}

	return max(queryMax(root.LeftChild, left, right), queryMax(root.RightChild, left, right))
}

type workProfit struct {
	d, p int
}

func Solution(difficulty []int, profit []int, worker []int) int {
	l := len(difficulty)
	dp := make([]workProfit, l)
	for i := 0; i < len(difficulty); i++ {
		dp[i] = workProfit{d: difficulty[i], p: profit[i]}
	}
	sort.Slice(dp, func(i, j int) bool {
		return dp[i].d < dp[j].d
	})

	tree := buildSegmentTree(dp, 0, l-1)
	ans := 0
	for _, n := range worker {
		if idx := sort.Search(l, func(i int) bool {
			return dp[i].d > n
		}); idx != 0 {
			ans += queryMax(tree, 0, idx-1)
		}
	}

	return ans
}
