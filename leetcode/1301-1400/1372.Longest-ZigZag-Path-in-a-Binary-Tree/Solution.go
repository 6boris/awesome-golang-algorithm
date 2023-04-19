package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode) int {
	ans := 1
	longestZigZag1(root, &ans)
	return ans - 1
}

func longestZigZag1(root *TreeNode, ans *int) (int, int) {
	if root == nil {
		return 0, 0
	}
	if root.Left == nil && root.Right == nil {
		return 1, 1
	}
	l, r := 1, 1
	rl, _ := longestZigZag1(root.Right, ans)
	_, lr := longestZigZag1(root.Left, ans)
	l += lr
	r += rl
	if l > *ans {
		*ans = l
	}
	if r > *ans {
		*ans = r
	}
	return l, r
}
