package Solution

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode) []*TreeNode {
	ans := make([]*TreeNode, 0)
	marshalTree(root, map[string]int{}, &ans)
	return ans
}

func marshalTree(root *TreeNode, count map[string]int, ans *[]*TreeNode) string {
	if root == nil {
		return "$"
	}
	now := fmt.Sprintf("%d,%s,%s", root.Val, marshalTree(root.Left, count, ans), marshalTree(root.Right, count, ans))
	count[now]++
	if count[now] == 2 {
		*ans = append(*ans, root)
	}
	return now
}
