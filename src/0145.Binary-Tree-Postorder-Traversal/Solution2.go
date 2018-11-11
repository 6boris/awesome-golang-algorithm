package Solution

type seqStack struct {
	data []*TreeNode
	tag  []int // 后续遍历准备
	top  int   // 数组下标
}

func postorderTraversal2(root *TreeNode) []int {
	var s seqStack
	s.top = -1
	result := []int{}
	if root != nil {
		for root != nil || s.top != -1 {
			for root != nil {
				s.top++
				s.data = append(s.data, root)
				s.tag = append(s.tag, 0)
				root = root.Left
			}

			if s.tag[s.top] == 0 {
				root = s.data[s.top]
				s.tag[s.top] = 1
				root = root.Right
			} else {
				for s.tag[s.top] == 1 {
					root = s.data[s.top]
					s.data = s.data[:len(s.data)-1]
					result = append(result, root.Val)
					s.top--
					if s.top < 0 {
						break
					}
				}
				root = nil
			}
		}
	}
	return result
}
