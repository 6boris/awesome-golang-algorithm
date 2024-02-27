package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode) []float64 {
	queue := []*TreeNode{root}
	ans := []float64{}
	for len(queue) > 0 {
		nq := []*TreeNode{}
		sum := 0
		for _, n := range queue {
			sum += n.Val
			if n.Left != nil {
				nq = append(nq, n.Left)
			}
			if n.Right != nil {
				nq = append(nq, n.Right)
			}
		}
		ans = append(ans, float64(sum)/float64(len(queue)))
		queue = nq
	}
	return ans
}
