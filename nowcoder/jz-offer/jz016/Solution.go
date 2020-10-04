package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	ans := &ListNode{}
	cur := ans
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val <= pHead2.Val {
			cur.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			cur.Next = pHead2
			pHead2 = pHead2.Next
		}
		cur = cur.Next
	}
	if pHead1 != nil {
		cur.Next = pHead1
	} else {
		cur.Next = pHead2
	}
	return ans.Next
}
func Merge2(pHead1 *ListNode, pHead2 *ListNode) *ListNode {

	if pHead1 == nil {
		return pHead2
	} else if pHead2 == nil {
		return pHead1
	}

	if pHead1.Val < pHead2.Val {
		pHead1.Next = Merge2(pHead1.Next, pHead2)
		return pHead1
	} else {
		pHead2.Next = Merge2(pHead1, pHead2.Next)
		return pHead2
	}
}
