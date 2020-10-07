package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return dfs(root, root)
}

func dfs(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val && dfs(t1.Right, t2.Left) && dfs(t1.Left, t2.Right)
}
