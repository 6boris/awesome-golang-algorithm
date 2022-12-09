package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root1 *TreeNode, root2 *TreeNode) bool {
	root1Leaves := leaves872(root1)
	root2Leaves := leaves872(root2)
	if len(root1Leaves) != len(root2Leaves) {
		return false
	}

	for i := 0; i < len(root1Leaves); i++ {
		if root1Leaves[i] != root2Leaves[i] {
			return false
		}
	}
	return true
}

func leaves872(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	left := leaves872(root.Left)
	right := leaves872(root.Right)

	ans := make([]int, 0)
	ans = append(ans, left...)
	ans = append(ans, right...)
	return ans
}
