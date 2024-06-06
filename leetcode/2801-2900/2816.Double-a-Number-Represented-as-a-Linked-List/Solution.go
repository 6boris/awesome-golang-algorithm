package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solution(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var do func(*ListNode, *int)
	do = func(cur *ListNode, cf *int) {
		if cur == nil {
			return
		}
		do(cur.Next, cf)
		v := cur.Val*2 + *cf
		*cf = v / 10
		v %= 10
		cur.Val = v
	}
	cf := 0
	do(head, &cf)
	if cf != 0 {
		node := &ListNode{Val: cf, Next: head}
		return node
	}
	return head
}
