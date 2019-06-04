package Solution

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p, q := headA, headB

	for p != q {
		if p != nil {
			p = p.Next
		} else {
			p = headB
		}

		if q != nil {
			q = q.Next
		} else {
			q = headA
		}
	}
	return p
}
