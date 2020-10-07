package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	} else if head.Val == val {
		return head.Next
	}

	ans := head
	for head.Next.Val != val {
		head = head.Next
	}
	head.Next = head.Next.Next
	return ans
}
