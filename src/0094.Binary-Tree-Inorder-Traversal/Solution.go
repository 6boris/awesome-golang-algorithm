package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	x := []int{}
	if root != nil {
		x = append(x, inorderTraversal(root.Left)...)
		x = append(x, root.Val)
		x = append(x, inorderTraversal(root.Right)...)
	}
	return x
}

const maxNum = 100

type seqStack struct {
	data [maxNum]*TreeNode
	tag  [maxNum]int // 后续遍历准备
	top  int         // 数组下标
}

func inorderTraversal2(root *TreeNode) []int {
	var s seqStack
	s.top = -1
	result := []int{}

	if root == nil {
		return result
	}
	for root != nil || s.top != -1 {
		for root != nil {
			s.top++
			s.data[s.top] = root
			root = root.Left
		}
		s.top--
		root = s.data[s.top+1]
		result = append(result, root.Val)
		root = root.Right
	}
	return result
}
