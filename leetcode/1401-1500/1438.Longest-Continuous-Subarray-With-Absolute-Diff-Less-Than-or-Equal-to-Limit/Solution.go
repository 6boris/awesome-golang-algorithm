package Solution

type SegmentTreeNode1438 struct {
	Left, Right int
	Min, Max    int
	LeftChild   *SegmentTreeNode1438
	RightChild  *SegmentTreeNode1438
}

func buildSegmentTree1438(arr []int, left, right int) *SegmentTreeNode1438 {
	if left == right {
		return &SegmentTreeNode1438{
			Left:  left,
			Right: right,
			Min:   arr[left],
			Max:   arr[left],
		}
	}

	mid := (left + right) / 2
	leftChild := buildSegmentTree1438(arr, left, mid)
	rightChild := buildSegmentTree1438(arr, mid+1, right)

	node := &SegmentTreeNode1438{
		Left:       left,
		Right:      right,
		Min:        min(leftChild.Min, rightChild.Min),
		Max:        max(leftChild.Max, rightChild.Max),
		LeftChild:  leftChild,
		RightChild: rightChild,
	}

	return node
}

func queryMin1438(node *SegmentTreeNode1438, left, right int) int {
	if node.Left == left && node.Right == right {
		return node.Min
	}

	mid := (node.Left + node.Right) / 2

	if right <= mid {
		return queryMin1438(node.LeftChild, left, right)
	} else if left > mid {
		return queryMin1438(node.RightChild, left, right)
	} else {
		leftMin := queryMin1438(node.LeftChild, left, mid)
		rightMin := queryMin1438(node.RightChild, mid+1, right)
		return min(leftMin, rightMin)
	}
}

func queryMax1438(node *SegmentTreeNode1438, left, right int) int {
	if node.Left == left && node.Right == right {
		return node.Max
	}

	mid := (node.Left + node.Right) / 2

	if right <= mid {
		return queryMax1438(node.LeftChild, left, right)
	} else if left > mid {
		return queryMax1438(node.RightChild, left, right)
	} else {
		leftMax := queryMax1438(node.LeftChild, left, mid)
		rightMax := queryMax1438(node.RightChild, mid+1, right)
		return max(leftMax, rightMax)
	}
}

// 滑动窗口, 线段树？
func Solution(nums []int, limit int) int {
	l := len(nums)
	tree := buildSegmentTree1438(nums, 0, l-1)
	start, end := 0, 0
	ans := 0

	var a, b int
	for ; end < l; end++ {
		a = queryMax1438(tree, start, end)
		b = queryMin1438(tree, start, end)
		if a-b <= limit {
			ans = max(ans, end-start+1)
			continue
		}
		for start < end {
			start++
			a = queryMax1438(tree, start, end)
			b = queryMin1438(tree, start, end)
			if a-b <= limit {
				break
			}
		}
	}
	return ans
}
