package Solution

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func Lowest(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}
	left := Lowest(root.left, p, q)
	right := Lowest(root.right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
