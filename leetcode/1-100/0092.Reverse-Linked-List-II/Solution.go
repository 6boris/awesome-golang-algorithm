package Solution

func Solution(head *ListNode, m, n int) *ListNode {
	newHead := new(ListNode)
	newHead.Next = head
	current := newHead

	for i := 0; i < m-1; i++ {
		current = current.Next
	}
	prev := current
	current = current.Next

	for i := 0; i < n-m; i++ {
		node := &ListNode{current.Next.Val, prev.Next}
		prev.Next = node
		current.Next = current.Next.Next
	}
	return newHead.Next
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	dummy := new(ListNode)
	head, dummy.Next = dummy, head

	for i := 0; i < m-1; i++ {
		head = head.Next
	}
	var curr, prev *ListNode = head.Next, nil
	for i := 0; i < n-m+1; i++ {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	head.Next.Next = curr
	head.Next = prev

	return dummy.Next
}

func reverseBetween2(head *ListNode, m int, n int) *ListNode {
	dummy := new(ListNode)
	head, dummy.Next = dummy, head
	for i := 0; i < m-1; i++ {
		head = head.Next
	}
	newHead, _ := reverseList(head.Next, n-m+1)
	head.Next = newHead
	// if m == 1 { return head.Next } else { return dummy.Next }
	return dummy.Next
}

// return new head and the head of the rest
func reverseList(head *ListNode, cnt int) (*ListNode, *ListNode) {
	if cnt == 1 {
		return head, head.Next
	}
	newHead, restHead := reverseList(head.Next, cnt-1)
	head.Next.Next = head
	head.Next = restHead
	return newHead, restHead
}
