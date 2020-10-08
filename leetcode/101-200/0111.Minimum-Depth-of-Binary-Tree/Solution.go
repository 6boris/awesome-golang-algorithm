package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//	左边为空只计算右边
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	//	右边为空只计算左边
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	//	正常计算左右2个子树
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
