package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

//func reverseList2(head *ListNode) *ListNode {
//	var pre  *ListNode
//	var next *ListNode
//	for head != nil {
//		next = head.Next
//		head.Next = pre
//		pre = head
//		head = next
//	}
//	return pre
//}

//func reverseList3(head *ListNode) *ListNode {
//	var prev *ListNode
//
//	for head != nil {
//		head.Next, head, prev = prev, head.Next, head
//	}
//	return prev
//}
