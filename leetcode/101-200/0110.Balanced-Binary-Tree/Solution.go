package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced_1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return max(dfs(node.Left), dfs(node.Right)) + 1
	}
	return (abs(dfs(root.Left)-dfs(root.Right)) <= 1) && isBalanced_1(root.Left) && isBalanced_1(root.Right)
}
func isBalanced_2(root *TreeNode) bool {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if abs(left-right) > 1 || left == -1 || right == -1 {
			return -1
		}
		return max(left, right) + 1
	}
	return dfs(root) >= 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
