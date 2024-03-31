package Solution

func Solution(root *TreeNode) {
	var inOrder func(**TreeNode, *TreeNode, []*TreeNode)
	inOrder = func(pre **TreeNode, tn *TreeNode, change []*TreeNode) {
		if tn == nil {
			return
		}
		inOrder(pre, tn.Left, change)
		if *pre != nil {
			if (*pre).Val > tn.Val {
				if change[0] != nil {
					change[1] = tn
				} else {
					change[0], change[1] = *pre, tn
				}

			}
		}
		*pre = tn
		inOrder(pre, tn.Right, change)
	}
	c := make([]*TreeNode, 2)
	var pre *TreeNode
	inOrder(&pre, root, c)
	c[0].Val, c[1].Val = c[1].Val, c[0].Val
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
