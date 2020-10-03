package Solution

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func Lowest(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if p.val < root.val && q.val < root.val {
		return Lowest(root.left, p, q)
	}
	if p.val > root.val && q.val > root.val {
		return Lowest(root.right, p, q)
	}
	return root
}
