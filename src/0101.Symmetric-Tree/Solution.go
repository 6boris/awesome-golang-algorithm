package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return root == nil || isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(left *TreeNode, right *TreeNode) bool {

	if left == nil || right == nil {
		return left == right
	}

	//	检查节点的值
	if left.Val != right.Val {
		return false
	}
	return isSymmetricHelper(left.Left, right.Right) &&
		isSymmetricHelper(left.Right, right.Left)
}
