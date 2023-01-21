package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := -1
	maxAncestorDiff1026(root, &ans)
	return ans
}

func check1026(ans *int, others []int) (int, int) {
	root, m, n := others[0], others[0], others[0]
	for idx := 1; idx < len(others); idx++ {
		diff := others[idx] - root
		if diff < 0 {
			diff = -diff
		}
		if *ans == -1 || diff > *ans {
			*ans = diff
		}
		if others[idx] > m {
			m = others[idx]
		}
		if others[idx] < n {
			n = others[idx]
		}
	}
	return m, n
}
func maxAncestorDiff1026(root *TreeNode, ans *int) (int, int) {
	other := []int{root.Val}
	if root.Left != nil {

		leftMax, leftMin := maxAncestorDiff1026(root.Left, ans)
		other = append(other, leftMax, leftMin)
	}

	if root.Right != nil {
		rightMax, rightMin := maxAncestorDiff1026(root.Right, ans)
		other = append(other, rightMax, rightMin)
	}

	return check1026(ans, other)
}
