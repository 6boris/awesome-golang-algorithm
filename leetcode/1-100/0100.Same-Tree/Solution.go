package Solution

// 递归遍历
// 时间复杂度 O(N)
// 时间复杂度 O(tree height)
func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}

	if p.Val == q.Val {
		return isSameTree(p.Left, q.Left) &&
			isSameTree(p.Right, q.Right)
	}

	return false
}
