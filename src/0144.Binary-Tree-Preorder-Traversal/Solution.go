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
