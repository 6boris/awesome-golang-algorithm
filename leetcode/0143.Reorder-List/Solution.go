package Solution

func Solution(head *ListNode) {
	if nil == head || nil == head.Next {
		return
	}
	mid := findMid(head)
	first, second := head, reverseList(mid.Next)
	mid.Next = nil

	newHead := new(ListNode)
	current := newHead
	for first != nil || second != nil {
		if first != nil {
			current.Next = &ListNode{first.Val, nil}
			current = current.Next
			first = first.Next
		}
		if second != nil {
			current.Next = &ListNode{second.Val, nil}
			current = current.Next
			second = second.Next
		}
	}
	head.Next = newHead.Next.Next
}

func findMid(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	newHead := head
	for head.Next != nil {
		next := head.Next
		head.Next = next.Next
		next.Next = newHead
		newHead = next
	}
	return newHead
}
