package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{Val: 0, Next: nil}
	n1, n2, tmp := l1, l2, node

	sum := 0

	for n1 != nil || n2 != nil {
		sum /= 10
		if n1 != nil {
			sum += n1.Val
			n1 = n1.Next

		}
		if n2 != nil {
			sum += n2.Val
			n2 = n2.Next
		}
		tmp.Next = &ListNode{Val: sum % 10}
		tmp = tmp.Next
	}

	if sum/10 != 0 {
		tmp.Next = &ListNode{Val: 1}
	}
	return node.Next

}
