package Solution

func Solution(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	head = &ListNode{Val: -1, Next: head}
	curr := head
	for curr != nil {
		psum := 0
		next := curr.Next
		for next != nil {
			psum += next.Val
			if psum == 0 {
				curr.Next = next.Next
			}
			next = next.Next
		}
		curr = curr.Next
	}
	return head.Next
}
