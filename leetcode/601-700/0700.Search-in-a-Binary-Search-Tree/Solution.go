package Solution

func Solution(root *TreeNode, val int) *TreeNode {
	walker := root
	for walker != nil {
		if walker.Val == val {
			return walker
		}
		if walker.Val < val {
			walker = walker.Right
			continue
		}
		walker = walker.Left
	}

	return nil
}
