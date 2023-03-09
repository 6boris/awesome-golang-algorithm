package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode) bool {
	if root == nil || root.Val&1 != 1 {
		return false
	}
	queue := []*TreeNode{root}
	next := 1
	for len(queue) > 0 {
		nextQ := make([]*TreeNode, 0)
		pre := -1

		for _, item := range queue {
			if item.Left != nil {

				if !((pre == -1 && item.Left.Val&1 != next) || (next&1 == 1 && item.Left.Val&1 == 0 && item.Left.Val < pre) || (next&1 == 0 && item.Left.Val&1 == 1 && item.Left.Val > pre)) {
					return false
				}
				pre = item.Left.Val
				nextQ = append(nextQ, item.Left)
			}
			if item.Right != nil {
				if !((pre == -1 && item.Right.Val&1 != next) || (next&1 == 1 && item.Right.Val&1 == 0 && item.Right.Val < pre) || (next&1 == 0 && item.Right.Val&1 == 1 && item.Right.Val > pre)) {
					return false
				}
				pre = item.Right.Val
				nextQ = append(nextQ, item.Right)
			}
		}

		next = 1 - next
		queue = nextQ
	}
	return true
}
