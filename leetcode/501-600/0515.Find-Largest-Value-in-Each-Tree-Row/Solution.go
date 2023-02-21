package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		nextQ := make([]*TreeNode, 0)
		max := 0
		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				nextQ = append(nextQ, queue[i].Left)
			}
			if queue[i].Right != nil {
				nextQ = append(nextQ, queue[i].Right)
			}
			if i == 0 {
				max = queue[i].Val
			} else if queue[i].Val > max {
				max = queue[i].Val
			}
		}
		ans = append(ans, max)
		queue = nextQ
	}
	return ans
}
