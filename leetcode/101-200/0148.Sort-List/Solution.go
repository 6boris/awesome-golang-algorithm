package Solution

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	nextHalfHead := slow.Next
	slow.Next = nil
	l1 := sortList(head)
	l2 := sortList(nextHalfHead)

	dummy := &ListNode{}
	cur := dummy
	for l1 != nil || l2 != nil {
		if l1 == nil || (l2 != nil && l1.Val >= l2.Val) {
			cur.Next = l2
			l2 = l2.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
		}

		cur = cur.Next
	}

	return dummy.Next
}
