package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}
	a := Solution(root1.Left, root2.Left)
	b := Solution(root1.Right, root2.Right)

	if a && b {
		return true
	}
	a = Solution(root1.Left, root2.Right)
	b = Solution(root1.Right, root2.Left)
	if a && b {
		return true
	}
	return false
}
