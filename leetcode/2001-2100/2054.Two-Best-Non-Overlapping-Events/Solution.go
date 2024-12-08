package Solution

import "sort"

type SegmentTreeNode2054 struct {
	Left, Right, Max      int
	LeftChild, RightChild *SegmentTreeNode2054
}

func buildSegmentTree2054(nums [][]int, left, right int) *SegmentTreeNode2054 {
	if left == right {
		return &SegmentTreeNode2054{Left: left, Right: right, Max: nums[left][2]}
	}

	mid := (left + right) / 2
	leftNode := buildSegmentTree2054(nums, left, mid)
	rightNode := buildSegmentTree2054(nums, mid+1, right)
	max := max(leftNode.Max, rightNode.Max)

	return &SegmentTreeNode2054{Left: left, Right: right, Max: max, LeftChild: leftNode, RightChild: rightNode}
}

func queryMax2054(root *SegmentTreeNode2054, left, right int) int {
	if root.Left >= left && root.Right <= right {
		return root.Max
	}

	if root.Right < left || root.Left > right {
		return -1
	}

	return max(queryMax2054(root.LeftChild, left, right), queryMax2054(root.RightChild, left, right))
}

func Solution(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		a, b := events[i], events[j]
		if a[0] == b[0] {
			return a[1] < b[1]
		}
		return a[0] < b[0]
	})
	l := len(events)

	tree := buildSegmentTree2054(events, 0, l-1)
	ans := 0
	for i, e := range events {
		ans = max(ans, e[2])
		idx := sort.Search(l-i-1, func(ii int) bool {
			return events[i+1+ii][0] > e[1]
		})
		if idx == l-i-1 {
			continue
		}
		ans = max(ans, e[2]+queryMax2054(tree, i+1+idx, l-1))
	}
	return ans
}
