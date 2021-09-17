package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd_1(head *ListNode, k int) *ListNode {
	c := 0
	for l := head; l != nil; l = l.Next {
		c++
	}
	klh := new(ListNode)
	for klh = head; c > k; klh = klh.Next {
		c--
	}
	return klh
}

func getKthFromEnd_2(head *ListNode, k int) *ListNode {
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
