package Solution

func Solution(root *Node) *Node {
	curr := root
	for curr != nil {
		if curr.Child == nil {
			curr = curr.Next
		} else {
			next := curr.Next
			node := Solution(curr.Child)
			curr.Child = nil
			curr.Next, node.Prev = node, curr
			for curr.Next != nil {
				curr = curr.Next
			}
			if next == nil {
				break
			}
			curr.Next, next.Prev = next, curr
			curr = curr.Next
		}
	}
	return root
}
