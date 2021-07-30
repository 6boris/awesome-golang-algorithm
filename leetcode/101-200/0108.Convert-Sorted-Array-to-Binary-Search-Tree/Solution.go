package Solution

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func Solution(nums []int) *TreeNode {
    if len(nums) == 0 { return nil }
    mid := len(nums) / 2
    root := &TreeNode{Val: nums[mid]}
    root.Left = Solution(nums[:mid])
    root.Right = Solution(nums[mid + 1:])
    return root
}