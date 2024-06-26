package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode) bool {
	if root == nil {
		return true
	}
	ok := true
	if root.Left != nil {
		if root.Val != root.Left.Val {
			return false
		}
		ok = ok && Solution(root.Left)
	}
	if root.Right != nil {
		if root.Val != root.Right.Val {
			return false
		}
		ok = ok && Solution(root.Right)
	}
	return ok
}
