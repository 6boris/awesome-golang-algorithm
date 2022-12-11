package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum_1(root *TreeNode) int {
	maxSum := -1 << 63
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(dfs(node.Left), 0)
		right := max(dfs(node.Right), 0)
		maxSum = max(maxSum, node.Val+left+right)

		return node.Val + max(left, right)
	}
	dfs(root)
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxPathSum_2(root *TreeNode) int {
	ans := 0
	ansInit := false
	maxPathSum124(root, &ans, &ansInit)
	return ans
}
func maxPathSum124(root *TreeNode, ans *int, ansInit *bool) (int, int) {
	if root == nil {
		return 0, 0
	}

	ll, lr := maxPathSum124(root.Left, ans, ansInit)
	rl, rr := maxPathSum124(root.Right, ans, ansInit)
	lmax := ll
	if lr > lmax {
		lmax = lr
	}

	rmax := rr
	if rl > rmax {
		rmax = rl
	}

	returnL, returnR := root.Val, root.Val
	l := lmax + root.Val
	r := rmax + root.Val
	if !*ansInit || l > *ans {
		*ans = l
		*ansInit = true
	}
	if !*ansInit || r > *ans {
		*ans = r
		*ansInit = true
	}
	if r := l + rmax; !*ansInit || r > *ans {
		*ans = r
		*ansInit = true
	}
	if !*ansInit || root.Val > *ans {
		*ans = root.Val
		*ansInit = true
	}
	if l > root.Val {
		returnL = l
	}
	if r > root.Val {
		returnR = r
	}
	return returnL, returnR
}
