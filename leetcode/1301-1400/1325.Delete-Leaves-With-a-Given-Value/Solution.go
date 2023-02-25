package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = Solution(root.Left, target)
	root.Right = Solution(root.Right, target)
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}
