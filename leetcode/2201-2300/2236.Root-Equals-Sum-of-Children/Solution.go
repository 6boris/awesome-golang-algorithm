package Solution

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

func Solution(root *TreeNode) bool {
	if root == nil {
		return false
	}
	sum := 0
	checkTree2236(root, &sum)
	return sum == root.Val*2
}

func checkTree2236(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	*sum += root.Val
	checkTree2236(root.Left, sum)
	checkTree2236(root.Right, sum)
}
