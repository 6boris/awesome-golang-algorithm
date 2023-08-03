package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root1 *TreeNode, root2 *TreeNode) []int {
	stack1 := make([]*TreeNode, 0)
	stack2 := make([]*TreeNode, 0)

	a, b := true, true
	ans := make([]int, 0)
	for len(stack1) > 0 || len(stack2) > 0 || root1 != nil || root2 != nil {
		if a {
			for ; root1 != nil; root1 = root1.Left {
				stack1 = append(stack1, root1)
			}
		}
		if b {
			for ; root2 != nil; root2 = root2.Left {
				stack2 = append(stack2, root2)
			}
		}
		l1, l2 := len(stack1), len(stack2)
		if l1 > 0 && l2 > 0 {
			x := stack1[l1-1]
			y := stack2[l2-1]
			if x.Val < y.Val {
				stack1 = stack1[:l1-1]
				ans = append(ans, x.Val)
				root1 = x.Right
				a, b = true, false
			} else {
				stack2 = stack2[:l2-1]
				ans = append(ans, y.Val)
				root2 = y.Right
				a, b = false, true
			}
			continue
		}

		if l1 > 0 {
			x := stack1[l1-1]
			ans = append(ans, x.Val)
			stack1 = stack1[:l1-1]
			root1 = x.Right
			a, b = true, false
			continue
		}
		if l2 > 0 {
			x := stack2[l2-1]
			stack2 = stack2[:l2-1]
			ans = append(ans, x.Val)
			root2 = x.Right
			a, b = false, true
		}
	}
	return ans
}
