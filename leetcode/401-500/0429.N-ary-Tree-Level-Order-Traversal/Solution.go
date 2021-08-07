package Solution

func Solution(root *Node) [][]int {
	r := make([][]int, 0)
	if root == nil {
		return r
	}
	level := 0
	queue := []*Node{root}
	for len(queue) > 0 {
		r = append(r, []int{})
		nextQueue := make([]*Node, 0)
		for _, node := range queue {
			r[level] = append(r[level], node.Val)
			if len(node.Children) > 0 {
				nextQueue = append(nextQueue, node.Children...)
			}
		}
		level++
		queue = nextQueue
	}

	return r
}
