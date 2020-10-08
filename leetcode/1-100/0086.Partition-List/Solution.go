package Solution

func Solution(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	before, after := &ListNode{}, &ListNode{}
	bh, ah := before, after
	for ; head != nil; head = head.Next {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
	}
	after.Next, before.Next = nil, ah.Next
	return bh.Next
}
