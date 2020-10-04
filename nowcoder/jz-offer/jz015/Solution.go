package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(pHead *ListNode) *ListNode {
	if pHead.Next == nil {
		return pHead
	}

	pre := pHead
	cur := pHead.Next
	next := pHead.Next.Next
	pre.Next = nil

	for cur != nil {
		next = cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}

	return pre
}
