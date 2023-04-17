package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode, k int) int {
	var inOrder func(*TreeNode)
	which := 1
	ans := -1
	inOrder = func(root *TreeNode) {
		if root == nil || ans != -1 {
			return
		}
		if root.Left != nil {
			inOrder(root.Left)
		}
		if which == k && ans == -1 {
			ans = root.Val
			return
		}
		which++
		if root.Right != nil {
			inOrder(root.Right)
		}
	}
	inOrder(root)
	return ans
}
