package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode) int {
	children, ans := 0, 0
	dfs2265(root, &children, &ans)
	return ans
}

func dfs2265(root *TreeNode, children, ans *int) int {
	if root == nil {
		return 0
	}

	lc, rc := 0, 0
	left := dfs2265(root.Left, &lc, ans)
	right := dfs2265(root.Right, &rc, ans)
	r := left + right + root.Val
	*children += lc + rc + 1
	if r/(lc+rc+1) == root.Val {
		*ans++
	}
	return r
}
