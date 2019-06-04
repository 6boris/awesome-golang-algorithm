package Solution

func Solution(head *ListNode, m, n int) *ListNode {
	newHead := new(ListNode)
	newHead.Next = head
	current := newHead

	for i := 0; i < m - 1; i++ {
		current = current.Next
	}
	prev := current
	current = current.Next

	for i := 0; i < n - m; i++ {
		node := &ListNode{current.Next.Val, prev.Next}
		prev.Next = node
		current.Next = current.Next.Next
	}
	return newHead.Next
}
