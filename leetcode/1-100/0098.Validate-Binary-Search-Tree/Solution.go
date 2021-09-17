package Solution

//	验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	INT_MAX := int(^uint(0) >> 1)
	INT_MIN := ^INT_MAX

	if root == nil {
		return true
	}

	return validate(root, INT_MAX, INT_MIN)
}

func validate(root *TreeNode, max, min int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min {
		return false
	}

	if root.Val >= max {
		return false
	}

	return validate(root.Left, root.Val, min) && validate(root.Right, max, root.Val)
}
