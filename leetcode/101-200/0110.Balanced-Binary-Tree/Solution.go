package Solution

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func Solution(root *TreeNode) bool {
    res, _ := check(root)
    return res
}

func check(root *TreeNode) (bool, int) {
    if root == nil {
        return true, 0
    }
    lbh, lh := check(root.Left)
    rbh, rh := check(root.Right)
    ishb := lbh && rbh && abs(lh - rh) <= 1
    return ishb, max(lh, rh) + 1
}

func abs(a int) int {
    if a < 0 { return -a }
    return a
}

func max(a, b int) int {
    if a > b { return a }
    return b
}
