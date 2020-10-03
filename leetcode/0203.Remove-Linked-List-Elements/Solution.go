package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	tmp := &ListNode{}

	tmp.Next = head

	for p := tmp; p != nil; {
		if p.Next != nil && p.Next.Val == val {
			p.Next = p.Next.Next

		} else {
			p = p.Next
		}
	}
	return tmp.Next
}
