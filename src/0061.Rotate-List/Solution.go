package Solution

func Solution(head *ListNode, k int) *ListNode {
    if nil == head || 0 == k {
		return head
	}
	head, length := cycleList(head)
	if k >= length {
		k = k % length
	}
	for i := 0; i < length - k - 1; i++ {
		head = head.Next
	}
	p := head.Next
	head.Next = nil
	return p
}

func cycleList(l *ListNode) (*ListNode, int) {
	head, length := l, 1
	for l.Next != nil {
		l = l.Next
		length++
	}
	l.Next = head
	return head, length
}