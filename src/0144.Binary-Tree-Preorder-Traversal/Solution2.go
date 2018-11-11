package Solution

type seqStack struct {
	data []*TreeNode
	tag  []int // 后续遍历准备
	top  int   // 数组下标
}

func preorderTraversal2(root *TreeNode) []int {
	var s seqStack
	s.top = -1
	result := []int{}

	if root == nil {
		return result
	}
	for root != nil || s.top != -1 {
		for root != nil {
			result = append(result, root.Val)
			s.top++
			s.data = append(s.data, root)
			root = root.Left
		}
		root = s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		root = root.Right
		s.top--
	}
	return result
}
