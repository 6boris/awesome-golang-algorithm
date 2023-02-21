package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode) int {
	ans, a := -1, -1
	inOrder783(root, &ans, &a)
	return ans
}

func inOrder783(root *TreeNode, ans, a *int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inOrder783(root.Left, ans, a)
	}
	if *a != -1 {
		diff := root.Val - *a
		if *ans == -1 || *ans > diff {
			*ans = diff
		}
	}
	*a = root.Val

	if root.Right != nil {
		inOrder783(root.Right, ans, a)
	}
}
