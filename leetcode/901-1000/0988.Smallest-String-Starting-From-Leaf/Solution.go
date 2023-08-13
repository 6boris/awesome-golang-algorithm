package Solution

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// a=97
func Solution(root *TreeNode) string {
	var dfs func(*TreeNode, string)
	ans := ""
	dfs = func(now *TreeNode, total string) {
		if now == nil {
			return
		}
		self := fmt.Sprintf("%c", now.Val+97) + total

		if now.Left == nil && now.Right == nil {
			if ans == "" || ans > self {
				ans = self
			}
			return
		}
		dfs(now.Left, self)
		dfs(now.Right, self)
	}
	dfs(root, "")
	return ans
}
