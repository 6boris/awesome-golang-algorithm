package Solution

func Solution(preorder []int) *TreeNode {
	return buildBstFromPreOrder(preorder)
}

func buildBstFromPreOrder(pre []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}

	v := pre[0]
	node := &TreeNode{Val: v}
	idx := 1
	for ; idx < len(pre); idx++ {
		if pre[idx] > v {
			break
		}
	}
	if idx > 1 {
		node.Left = buildBstFromPreOrder(pre[1:idx])
	}

	if idx < len(pre) {
		node.Right = buildBstFromPreOrder(pre[idx:])
	}

	return node
}
