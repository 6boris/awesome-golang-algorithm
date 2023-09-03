package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val < val {
		return &TreeNode{Val: val, Left: root}
	}
	r := root.Right
	if r == nil {
		root.Right = &TreeNode{Val: val}
		return root
	}
	root.Right = Solution(root.Right, val)
	return root
}
