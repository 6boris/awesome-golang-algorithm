package Solution


func pathSum(root *TreeNode, sum int) [][]int {
	var slice [][]int
	slice = pathSumHelper(root, sum, slice, []int(nil))
	return slice
}

func pathSumHelper(n *TreeNode, sum int, slice [][]int, stack []int) [][]int {
	if n == nil {
		return slice
	}
	sum -= n.Val
	stack = append(stack, n.Val)

	if sum == 0 && n.Left == nil && n.Right == nil {
		slice = append(slice, append([]int(nil), stack...))
		stack = stack[:len(stack)-1]
	}

	slice = pathSumHelper(n.Left, sum, slice, stack)
	slice = pathSumHelper(n.Right, sum, slice, stack)
	return slice
}
