package Solution

//	递归遍历
//	时间复杂度 O(N)
//	时间复杂度 O(tree height)
func isSameTree(p *TreeNode, q *TreeNode) bool {
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

//	递归遍历
//	时间复杂度 O(N)
//	时间复杂度 O(tree height)

func isSameTree2(p *TreeNode, q *TreeNode) bool {
	qP := []*TreeNode{p}
	qQ := []*TreeNode{q}

	for len(qP) != 0 && len(qQ) != 0 {
		pNode := qP[0]
		qP = qP[1:]

		qNode := qQ[0]
		qQ = qQ[1:]

		if pNode == nil && qNode == nil {
			continue
		}
		if pNode == nil && qNode != nil || pNode != nil && qNode == nil {
			return false
		}
		if pNode.Val != qNode.Val {
			return false
		}

		qP = append(qP, pNode.Left, pNode.Right)
		qQ = append(qQ, qNode.Left, qNode.Right)
	}

	if len(qP) == 0 && len(qQ) == 0 {
		return true
	}
	return false
}
