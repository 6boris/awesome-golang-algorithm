package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode) *TreeNode {
	sum := 0
	var preOrder func(*TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		sum += root.Val
		if root.Left != nil {
			preOrder(root.Left)
		}
		if root.Right != nil {
			preOrder(root.Right)
		}
	}
	preOrder(root)
	/*
		tmp := sum
		for i := 0; i < len(nodes); i++ {
			tmp -= nodes[i]
			nodes[i] += tmp
		}
	*/
	tmpSum := 0
	var inOrder func(*TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			inOrder(root.Left)
		}
		tmpSum += root.Val
		root.Val += sum - tmpSum
		if root.Right != nil {
			inOrder(root.Right)
		}
	}
	inOrder(root)
	return root
}
