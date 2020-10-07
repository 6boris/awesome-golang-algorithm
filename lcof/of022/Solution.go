package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	left, right := head, head
	for i := 0; i < k; i++ {
		right = right.Next
	}
	for right != nil {
		right = right.Next
		left = left.Next
	}
	return left
}
