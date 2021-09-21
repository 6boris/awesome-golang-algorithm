package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree_1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree_1(root.Left)
	right := invertTree_1(root.Right)
	root.Left = right
	root.Right = left

	return root
}
