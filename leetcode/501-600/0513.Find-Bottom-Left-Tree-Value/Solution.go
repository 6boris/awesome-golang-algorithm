package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solution(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	var ans *TreeNode
	for len(queue) > 0 {
		nextQ := make([]*TreeNode, 0)
		ans = queue[0]
		for _, n := range queue {
			if n.Left != nil {
				nextQ = append(nextQ, n.Left)
			}
			if n.Right != nil {
				nextQ = append(nextQ, n.Right)
			}
		}
		queue = nextQ
	}
	return ans.Val
}
