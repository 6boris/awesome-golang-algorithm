package Solution

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST_1(root *TreeNode) bool {
	var dfs func(*TreeNode, int, int) bool
	dfs = func(node *TreeNode, low, high int) bool {
		if node == nil {
			return true
		}
		if node.Val <= low || node.Val >= high {
			return false
		}
		return dfs(node.Left, low, node.Val) && dfs(node.Right, node.Val, high)
	}
	return dfs(root, math.MinInt64, math.MaxInt64)
}

func isValidBST_2(root *TreeNode) bool {
	tmp, stk := math.MinInt64, []*TreeNode{}
	node := root
	for node != nil || len(stk) > 0 {
		for node != nil {
			stk = append(stk, node)
			node = node.Left
		}
		node = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if node.Val <= tmp {
			return false
		}
		tmp = node.Val
		node = node.Right
	}
	return true
}
