package Solution

func Solution(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	node := &TreeNode{}
	mid := len(nums) / 2
	node.Val = nums[mid]
	if mid > 0 {
		node.Left = Solution(nums[:mid])
	}
	if mid < len(nums)-1 {
		node.Right = Solution(nums[mid+1:])
	}

	return node
}
