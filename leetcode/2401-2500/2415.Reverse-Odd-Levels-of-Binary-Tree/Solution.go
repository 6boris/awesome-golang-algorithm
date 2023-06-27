package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode) *TreeNode {
	q := []*TreeNode{root}
	odd := false
	alloc := 2
	for len(q) > 0 {
		next := make([]*TreeNode, alloc)
		start, end := 0, alloc-1
		allNil := false

		for s, e := 0, len(q)-1; s <= e; s, e = s+1, e-1 {
			if odd {
				q[s].Val, q[e].Val = q[e].Val, q[s].Val
			}
			if q[s].Left == nil || q[s].Right == nil {
				allNil = true
				continue
			}

			next[start] = q[s].Left
			next[start+1] = q[s].Right
			next[end] = q[e].Right
			next[end-1] = q[e].Left
			start += 2
			end -= 2

		}

		if allNil {
			break
		}
		odd = !odd
		alloc *= 2
		q = next
	}
	return root
}

func Solution1(root *TreeNode) *TreeNode {
	reverseOddLevels12(root.Left, root.Right, true)
	return root
}
func reverseOddLevels12(left, right *TreeNode, odd bool) {
	if left == nil || right == nil {
		return
	}
	if odd {
		left.Val, right.Val = right.Val, left.Val
	}
	reverseOddLevels12(left.Left, right.Right, !odd)
	reverseOddLevels12(left.Right, right.Left, !odd)
}
