package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Val == 0 {
		return false
	}
	if root.Val == 1 {
		return true
	}
	l := Solution(root.Left)
	r := Solution(root.Right)
	if root.Val == 2 {
		return l || r
	}
	return l && r
}
