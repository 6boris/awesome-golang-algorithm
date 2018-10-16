package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fastNode := head
	slowNode := head
	for n > 0 {
		fastNode = fastNode.Next
		n--
	}

	if fastNode != nil {
		for fastNode.Next != nil {
			fastNode = fastNode.Next
			slowNode = slowNode.Next
		}
		slowNode.Next = slowNode.Next.Next

	} else {
		head = head.Next
	}
	return head
}
