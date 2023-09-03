package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode, limit int) *TreeNode {
	var dfs func(*TreeNode, int) *TreeNode
	dfs = func(tree *TreeNode, sum int) *TreeNode {
		if tree == nil {
			return nil
		}
		s := sum + tree.Val
		if tree.Left == nil && tree.Right == nil {
			if s < limit {
				return nil
			}
			return tree
		}
		tree.Left = dfs(tree.Left, s)
		tree.Right = dfs(tree.Right, s)
		if tree.Left == nil && tree.Right == nil {
			return nil
		}
		return tree
	}
	return dfs(root, 0)
}
