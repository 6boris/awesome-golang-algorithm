package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

// 循环
func swapPairs(head *ListNode) *ListNode {
	preHead := &ListNode{Val: 0, Next: nil}
	cur := preHead
	preHead.Next = head
	for cur.Next != nil && cur.Next.Next != nil {
		tmp := cur.Next.Next
		cur.Next.Next = tmp.Next
		tmp.Next = cur.Next
		cur.Next = tmp
		cur = cur.Next.Next
	}
	return preHead.Next
}

//	递归
func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := head.Next
	head.Next = swapPairs(node.Next)
	node.Next = head
	return node
}
