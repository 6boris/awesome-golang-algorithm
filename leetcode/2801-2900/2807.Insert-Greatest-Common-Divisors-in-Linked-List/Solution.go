package Solution

func gcd2807(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solution(head *ListNode) *ListNode {
	a, b := head, head.Next
	for b != nil {
		node := &ListNode{Val: gcd2807(a.Val, b.Val)}
		a.Next = node
		node.Next = b
		a, b = b, b.Next
	}
	return head
}
