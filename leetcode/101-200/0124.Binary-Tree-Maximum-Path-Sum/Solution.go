package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum_1(root *TreeNode) int {
	maxSum := -1 << 63
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(dfs(node.Left), 0)
		right := max(dfs(node.Right), 0)
		maxSum = max(maxSum, node.Val+left+right)

		return node.Val + max(left, right)
	}
	dfs(root)
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
