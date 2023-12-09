package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode) int {
	steps := 0
	var dfs func(*TreeNode) int
	dfs = func(tree *TreeNode) int {
		if tree == nil {
			return 0
		}
		left := dfs(tree.Left)
		right := dfs(tree.Right)
		next := tree.Val + left + right - 1

		if left < 0 {
			left = -left
		}
		steps += left
		if right < 0 {
			right = -right
		}
		steps += right
		return next
	}
	dfs(root)
	return steps
}
