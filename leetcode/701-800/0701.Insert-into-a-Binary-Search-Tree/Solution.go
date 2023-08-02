package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val < val {
		root.Right = Solution(root.Right, val)
	} else if root.Val > val {
		root.Left = Solution(root.Left, val)
	}
	return root
}
