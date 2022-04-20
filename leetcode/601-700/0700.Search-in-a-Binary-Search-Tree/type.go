package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isSameTree(root, cmp *TreeNode) bool {
	if root == nil && cmp == nil {
		return true
	}

	if root == nil || cmp == nil {
		return false
	}

	if root.Val != cmp.Val {
		return false
	}

	return isSameTree(root.Left, cmp.Left) && isSameTree(root.Right, cmp.Right)
}
