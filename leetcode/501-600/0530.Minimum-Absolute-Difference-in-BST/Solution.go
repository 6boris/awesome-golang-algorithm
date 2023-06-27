package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode) int {
	target := -1
	ans := -1
	dfs530(root, &target, &ans)
	return ans
}

func dfs530(root *TreeNode, target, ans *int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		dfs530(root.Left, target, ans)
	}
	if *target != -1 {
		diff := root.Val - *target
		if *ans == -1 || diff < *ans {
			*ans = diff
		}
	}
	*target = root.Val
	if root.Right != nil {
		dfs530(root.Right, target, ans)
	}
}
