package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode) []int {
	freq := -1
	pre := 1
	first := true
	maxCount := 0
	ans := make([]int, 0)
	var inOrder func(tree *TreeNode)
	inOrder = func(tree *TreeNode) {
		if tree == nil {
			return
		}

		inOrder(tree.Left)
		if first {
			freq = 1
			pre = tree.Val
			first = false
		} else {
			if tree.Val == pre {
				freq++
			} else {
				x := pre
				pre = tree.Val
				if freq == maxCount {
					ans = append(ans, x)
				} else if freq > maxCount {
					maxCount = freq
					ans = []int{x}
				}
				freq = 1
			}
		}
		inOrder(tree.Right)
	}
	inOrder(root)
	if freq == maxCount {
		ans = append(ans, pre)
	} else if freq > maxCount {
		ans = []int{pre}
	}
	return ans
}
