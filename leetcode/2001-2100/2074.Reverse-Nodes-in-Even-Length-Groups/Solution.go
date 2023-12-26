package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func rev2074(head, headCopy **ListNode, tail *ListNode, c *int, l int) {
	if *headCopy != tail {
		rev2074(head, &(*headCopy).Next, tail, c, l)
	}
	if *c == l/2 {
		return
	}
	(*head).Val, (*headCopy).Val = (*headCopy).Val, (*head).Val
	*head = (*head).Next
	*c++
}

func Solution(head *ListNode) *ListNode {
	var (
		walker = head
		pre    *ListNode
	)

	l, cur := 1, 0
	for ; walker != nil; walker = walker.Next {
		cur++
		if cur != l && walker.Next != nil {
			continue
		}

		if cur&1 == 0 {
			c := 0
			source := pre.Next
			rev2074(&source, &source, walker, &c, cur)
		}
		pre = walker
		cur = 0
		l++
	}
	return head
}
