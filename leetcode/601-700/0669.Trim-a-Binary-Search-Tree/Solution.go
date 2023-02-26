package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return Solution(root.Right, low, high)
	}
	if root.Val > high {
		return Solution(root.Left, low, high)
	}

	root.Left = Solution(root.Left, low, high)
	root.Right = Solution(root.Right, low, high)
	return root
}
