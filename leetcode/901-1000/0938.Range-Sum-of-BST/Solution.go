package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode, low, high int) int {
	var inOrder func(*TreeNode, *bool, *bool)
	ans := 0
	inOrder = func(root *TreeNode, start, end *bool) {
		if root == nil {
			return
		}
		if root.Left != nil {
			inOrder(root.Left, start, end)
		}
		if root.Val == low {
			*start = true
		}
		if high == root.Val {
			*end = true
			ans += root.Val
		}
		if *start && !*end {
			ans += root.Val
		}
		if !*end && root.Right != nil {
			inOrder(root.Right, start, end)
		}
	}
	start, end := false, false
	inOrder(root, &start, &end)
	return ans
}
