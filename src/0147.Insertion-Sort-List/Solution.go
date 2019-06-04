package Solution

func insertionSortList(head *ListNode) *ListNode {
	dummy := new(ListNode)
	for head != nil {
		next, p := head.Next, dummy

		for p.Next != nil && p.Next.Val <= head.Val {
			p = p.Next
		}

		head.Next = p.Next
		p.Next = head
		head = next
	}
	return dummy.Next
}
