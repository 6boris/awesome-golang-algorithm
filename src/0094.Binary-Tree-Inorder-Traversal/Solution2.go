package Solution

type seqStack struct {
	data []*TreeNode
	top  int // 数组下标
}

func inorderTraversal2(root *TreeNode) []int {
	var s seqStack
	s.top = -1
	result := []int{}

	if root != nil {
		for root != nil || s.top != -1 {
			for root != nil {
				s.top++
				s.data = append(s.data, root)
				root = root.Left
			}
			s.top--
			root = s.data[len(s.data)-1]
			s.data = s.data[:len(s.data)-1]
			result = append(result, root.Val)
			root = root.Right
		}
	}
	return result
}
