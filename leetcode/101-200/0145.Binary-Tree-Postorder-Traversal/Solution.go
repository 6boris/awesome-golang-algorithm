package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal_1(root *TreeNode) []int {
	ans := make([]int, 0)
	if root != nil {
		ans = append(ans, postorderTraversal_1(root.Left)...)
		ans = append(ans, postorderTraversal_1(root.Right)...)
		ans = append(ans, root.Val)
	}
	return ans
}

func postorderTraversal_2(root *TreeNode) []int {
	stk, ans := []*TreeNode{}, make([]int, 0)
	var prev *TreeNode
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.Right == nil || root.Right == prev {
			ans = append(ans, root.Val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root)
			root = root.Right
		}
	}
	return ans
}
