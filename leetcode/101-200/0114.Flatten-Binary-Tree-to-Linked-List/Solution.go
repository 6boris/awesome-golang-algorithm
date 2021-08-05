package Solution

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func Solution(root *TreeNode)  {
    if root == nil { return }
    Solution(root.Left)
    if root.Left != nil {
        left, right := root.Left, root.Right
        root.Right = left
        for left.Right != nil {
            left = left.Right
        }
        left.Right = right
        root.Left = nil
    }
    Solution(root.Right)    
}