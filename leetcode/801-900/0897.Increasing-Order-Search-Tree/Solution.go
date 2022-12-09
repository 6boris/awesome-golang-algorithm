package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var tree, walker *TreeNode
	var inOrder func(*TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			inOrder(root.Left)
		}

		n := &TreeNode{Val: root.Val}
		if tree == nil {
			tree = n
			walker = n
		} else {
			walker.Right = n
			walker = walker.Right
		}
		if root.Right != nil {
			inOrder(root.Right)
		}
	}
	inOrder(root)
	return tree
}
