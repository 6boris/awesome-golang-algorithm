package Solution


func Solution(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}
	newHead := new(ListNode)
	newHead.Next = &ListNode{head.Val, nil}

	head = head.Next
	for head != nil {
		prev, current := newHead, newHead.Next
		for current != nil {
			if head.Val < current.Val {
				prev.Next = &ListNode{head.Val, current}
				break
			}
			prev, current = prev.Next, current.Next
		}
		if nil == current && head.Val >= prev.Val {
			prev.Next = &ListNode{head.Val, nil}
		}
		head = head.Next
	}
	return newHead.Next
}

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

