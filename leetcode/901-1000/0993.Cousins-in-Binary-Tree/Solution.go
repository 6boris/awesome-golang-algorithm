package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode, x int, y int) bool {
	var (
		xnode  *TreeNode
		ynode  *TreeNode
		xd, yd int
	)
	var dfs func(*TreeNode, *TreeNode, int)
	dfs = func(fa, tree *TreeNode, dep int) {
		if tree == nil {
			return
		}
		if tree.Val == x {
			xnode = fa
			xd = dep
		}
		if tree.Val == y {
			ynode = fa
			yd = dep
		}
		dfs(tree, tree.Left, dep+1)
		dfs(tree, tree.Right, dep+1)
	}
	dfs(nil, root, 0)
	if xd != yd {
		return false
	}
	return xnode != ynode
}
