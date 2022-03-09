package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tmp := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 != nil {
		tmp.Next = l1
	} else {
		tmp.Next = l2
	}

	return head.Next
}

func mergeTwoLists1(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	n := &ListNode{}
	if l1.Val < l2.Val {
		n.Val = l1.Val
		n.Next = mergeTwoLists1(l1.Next, l2)
	} else {
		n.Val = l2.Val
		n.Next = mergeTwoLists1(l1, l2.Next)
	}
	return n
}
