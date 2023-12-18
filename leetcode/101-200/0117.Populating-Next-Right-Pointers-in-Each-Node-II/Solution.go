package Solution

type Node1 struct {
	Val               int
	Left, Right, Next *Node1
}

func Solution(root *Node1) *Node1 {
	if root == nil {
		return nil
	}
	queue := []*Node1{root}
	for len(queue) > 0 {
		nq := make([]*Node1, 0)
		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				nq = append(nq, queue[i].Left)
			}
			if queue[i].Right != nil {
				nq = append(nq, queue[i].Right)
			}

			if i != len(queue)-1 {
				queue[i].Next = queue[i+1]
			}
		}
		queue = nq
	}
	return root
}
