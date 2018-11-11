package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	x := []int{}
	if root != nil {
		x = append(x, inorderTraversal(root.Left)...)
		x = append(x, root.Val)
		x = append(x, inorderTraversal(root.Right)...)
	}
	return x
}
