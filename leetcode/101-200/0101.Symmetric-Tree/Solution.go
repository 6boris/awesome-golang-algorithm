package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//	递归
func isSymmetric(root *TreeNode) bool {
	return root == nil || isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	//	左右节点一个不存在
	if left == nil || right == nil {
		return left == right
	}

	//	检查节点的值
	if left.Val != right.Val {
		return false
	}
	return isSymmetricHelper(left.Left, right.Right) &&
		isSymmetricHelper(left.Right, right.Left)
}

//	循环
func isSymmetric_BFS(root *TreeNode) bool {
	queue, left, right := []*TreeNode{root.Left, root.Right}, &TreeNode{}, &TreeNode{}

	for len(queue) > 0 {
		left, right, queue = queue[0], queue[1], []*TreeNode{}

		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}

		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}
	return true
}

func isSymmetric_DFS(root *TreeNode) bool {
	stack, left, right := []*TreeNode{root.Left, root.Right}, &TreeNode{}, &TreeNode{}
	for len(stack) > 0 {
		left, right = stack[len(stack)-1], stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil {
			return false
		}

		if left.Val != right.Val {
			return false
		}

		stack = append(stack, left.Left, right.Right, left.Right, right.Left)
	}
	return true
}
