package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(descriptions [][]int) *TreeNode {
	indies := [100001]*TreeNode{}
	children := [100001]bool{}
	for _, desc := range descriptions {
		parent, child, left := desc[0], desc[1], desc[2]
		children[child] = true
		if indies[parent] == nil {
			indies[parent] = &TreeNode{Val: parent}
		}
		if indies[child] == nil {
			indies[child] = &TreeNode{Val: child}
		}
		if left == 1 {
			indies[parent].Left = indies[child]
		} else {
			indies[parent].Right = indies[child]
		}
	}
	for _, desc := range descriptions {
		if !children[desc[0]] {
			return indies[desc[0]]
		}
	}

	return nil
}
