package Solution

type Node struct {
	Val   int
	Next  *Node
	Prev  *Node
	Child *Node
}

func MakeNode(nodes []int) *Node {
	if nodes == nil || len(nodes) < 1 {
		return &Node{}
	}
	node := &Node{}
	head := node
	node.Val = nodes[0]
	for i := 1; i < len(nodes); i++ {
		node.Next = &Node{Val: nodes[i], Prev: node}
		node = node.Next
	}
	return head
}

func Check(headA, headB *Node) bool {
	if headA == nil && headB == nil {
		return true
	}
	for curA, curB := headA, headB; curA != nil && curB != nil; curA, curB = curA.Next, curB.Next {
		if curA.Val != curB.Val {
			return false
		}
		// if curA.Child != nil || curB.Child != nil { return false }
	}
	return true
}
