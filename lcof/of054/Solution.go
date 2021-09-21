package Solution

func Solution(x bool) bool {
	return x
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest_1(root *TreeNode, k int) int {
	ans := -1
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		if k == 0 {
			return
		}
		k--
		if k == 0 {
			ans = node.Val
		}
		dfs(node.Left)

	}
	dfs(root)
	return ans
}
