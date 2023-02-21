package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(nums []int) *TreeNode {
	return b(nums, 0, len(nums)-1)
}

func b(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	targetIndex := left
	for i := left + 1; i <= right; i++ {
		if nums[i] > nums[targetIndex] {
			targetIndex = i
		}
	}
	node := &TreeNode{
		Val: nums[targetIndex],
	}
	node.Left = b(nums, left, targetIndex-1)
	node.Right = b(nums, targetIndex+1, right)
	return node
}
