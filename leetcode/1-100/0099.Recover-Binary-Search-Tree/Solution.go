package Solution

func Solution(x bool) bool {
	return x
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (res []int) {
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return
}
