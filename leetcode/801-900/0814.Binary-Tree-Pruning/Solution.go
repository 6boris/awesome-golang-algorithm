package Solution

func Solution(root *TreeNode) *TreeNode {
	newTree, del := pruneTreeAux(root)
	if del {
		return nil
	}
	return newTree
}

func pruneTreeAux(root *TreeNode) (*TreeNode, bool) {
	if root == nil {
		return nil, true
	}

	leftNode, delLeft := pruneTreeAux(root.Left)
	rightNode, delRight := pruneTreeAux(root.Right)
	delSelf := true
	n := &TreeNode{}
	if !delLeft {
		delSelf = false
		n.Left = leftNode
	}
	if !delRight {
		delSelf = false
		n.Right = rightNode
	}

	delSelf = root.Val == 0 && delSelf
	if delSelf {
		return nil, delSelf
	}
	n.Val = root.Val
	n.Left, n.Right = leftNode, rightNode
	return n, delSelf
}
