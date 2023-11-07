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

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	m := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		m[tmp] = true
	}

	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if _, ok := m[tmp]; ok {
			return tmp
		}
	}

	return nil
}
