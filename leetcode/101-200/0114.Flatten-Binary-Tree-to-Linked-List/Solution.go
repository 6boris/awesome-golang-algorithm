package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solution(root *TreeNode) {
	if root == nil {
		return
	}
	Solution(root.Left)
	if root.Left != nil {
		left, right := root.Left, root.Right
		root.Right = left
		for left.Right != nil {
			left = left.Right
		}
		left.Right = right
		root.Left = nil
	}
	Solution(root.Right)
}


func flatten114(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	end := root
	rightStart, rightEnd := flatten114(root.Right)
	leftStart, leftEnd := flatten114(root.Left)
	if leftStart != nil {
		end.Right = leftStart
		end = leftEnd
	}
	if rightStart != nil {
		end.Right = rightStart
		end = rightEnd
	}
	root.Left = nil
	return root, end
}
func Solution2(root *TreeNode) {
	flatten114(root)
	return
}
