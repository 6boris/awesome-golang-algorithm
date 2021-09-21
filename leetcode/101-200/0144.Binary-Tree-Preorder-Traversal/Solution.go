package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal_1(root *TreeNode) []int {
	ans := make([]int, 0)
	if root != nil {
		ans = append(ans, root.Val)
		ans = append(ans, preorderTraversal_1(root.Left)...)
		ans = append(ans, preorderTraversal_1(root.Right)...)
	}
	return ans
}
func preorderTraversal_2(root *TreeNode) []int {
	ans, stk := make([]int, 0), []*TreeNode{}
	node := root
	for node != nil || len(stk) > 0 {
		for node != nil {
			ans = append(ans, node.Val)
			stk = append(stk, node)
			node = node.Left
		}
		node = stk[len(stk)-1].Right
		stk = stk[:len(stk)-1]
	}
	return ans
}
