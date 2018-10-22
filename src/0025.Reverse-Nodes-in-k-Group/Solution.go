package Solution

func reverseKGroup(head *ListNode, k int) *ListNode {
	curr := head
	count := 0
	for curr != nil && count != k {
		curr = curr.Next
		count++
	}

	if count == k {
		curr = reverseKGroup(curr, k)
		for count > 0 {
			tmp := head.Next
			head.Next = curr
			curr = head
			head = tmp
			count--
		}
		head = curr
	}
	return head
}
