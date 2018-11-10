package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var stack []*TreeNode
	var nodes []int

	stack = append(stack, root)

	for len(stack) != 0 {
		lastNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nodes = append(nodes, lastNode.Val)

		if lastNode.Right != nil {
			stack = append(stack, lastNode.Right)
		}
		if lastNode.Left != nil {
			stack = append(stack, lastNode.Left)
		}
	}

	return nodes
}

const maxNum = 100

type seqStack struct {
	data [maxNum]*TreeNode
	tag  [maxNum]int // 后续遍历准备
	top  int         // 数组下标
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
			s.data[s.top] = root
			root = root.Left
		}
		s.top--
		root = s.data[s.top+1]
		root = root.Right
	}
	return result
}
