package Solution

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	curr := head

	for curr.Next != nil {
		if curr.Next.Val == curr.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
