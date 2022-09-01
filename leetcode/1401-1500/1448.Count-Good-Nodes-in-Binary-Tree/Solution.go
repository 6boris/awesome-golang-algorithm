package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solution(root *TreeNode) int {
	ans := 0
	goodNodesDfs(root, &ans, root.Val)
	return ans
}

func goodNodesDfs(root *TreeNode, ans *int, cmp int) {
	if root == nil {
		return
	}
	if root.Val >= cmp {
		*ans++
		cmp = root.Val
	}
	goodNodesDfs(root.Left, ans, cmp)
	goodNodesDfs(root.Right, ans, cmp)
}
