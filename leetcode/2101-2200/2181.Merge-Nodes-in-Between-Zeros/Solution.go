package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solution(head *ListNode) *ListNode {
	skipNode := head
	walker := head.Next
	for ; walker.Next != nil; walker = walker.Next {
		if walker.Val == 0 {
			skipNode.Next = walker
			skipNode = walker
			continue
		}
		skipNode.Val += walker.Val
	}
	skipNode.Next = nil
	return head
}
