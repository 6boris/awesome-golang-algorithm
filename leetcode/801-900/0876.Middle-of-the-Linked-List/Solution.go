package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solution(head *ListNode) *ListNode {
	tort, hare := head, head
	for hare != nil && hare.Next != nil {
		tort = tort.Next
		hare = hare.Next.Next
	}
	return tort
}
