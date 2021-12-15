package Solution

func findTiltAux(root *TreeNode, tilt *int) int {
	if root == nil {
		return 0
	}

	left := findTiltAux(root.Left, tilt)
	right := findTiltAux(root.Right, tilt)
	abs := left - right
	if abs < 0 {
		abs = -abs
	}
	*tilt += abs
	return left + right + root.Val
}

func Solution(root *TreeNode) int {
	tilt := 0
	_ = findTiltAux(root, &tilt)
	return tilt
}
