package Solution

import "math"

func Solution(root *TreeNode) int {
	res := math.MinInt32
	helper(root, &res)
	return res
}

func helper(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left, res)
	right := helper(root.Right, res)
	*res = max(*res, root.Val, left+root.Val, root.Val+right, left+root.Val+right)
	return max(left+root.Val, root.Val+right, root.Val)
}

func max(nums ...int) int {
	maxVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	return maxVal
}
