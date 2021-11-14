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

//func reverseList_2(head *ListNode) *ListNode {
//	var prev *ListNode
//	curr := head
//	for curr != nil {
//		next := curr.Next
//		curr.Next = prev
//		prev = curr
//		curr = next
//	}
//	return prev
//}

//func reverseList_3(head *ListNode) *ListNode {
//	var prev *ListNode
//
//	for head != nil {
//		head.Next, head, prev = prev, head.Next, head
//	}
//	return prev
//}
