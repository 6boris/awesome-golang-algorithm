package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solution(head *ListNode, k int) *ListNode {
	l := 0
	for walker := head; walker != nil; walker = walker.Next {
		l++
	}
	if k > l/2 {
		k = l - k + 1
	}
	steps := l - k*2 + 1
	fromHead := head
	for i := 1; i < k; i++ {
		fromHead = fromHead.Next
	}
	fromEnd := fromHead
	for ; steps > 0; steps-- {
		fromEnd = fromEnd.Next
	}
	fromHead.Val, fromEnd.Val = fromEnd.Val, fromHead.Val

	return head
}

func Solution1(head *ListNode, k int) *ListNode {
	walker := head
	for i := 1; i < k; i++ {
		walker = walker.Next
	}
	one := walker
	two := head
	for ; walker.Next != nil; walker = walker.Next {
		two = two.Next
	}
	one.Val, two.Val = two.Val, one.Val
	return head
}
