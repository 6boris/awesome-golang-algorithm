package Solution

func Solution(x bool) bool {
	return x
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

//
//func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
//	head := &ListNode{}
//	newList := head
//
//	for l1 != nil || l2 != nil {
//		if l1 == nil || (l2 != nil && l1.Val > l2.Val) {
//			newList.Next = l2
//			newList = l2
//			l2 = l2.Next
//		} else {
//			newList.Next = l1
//			newList = l1
//			l1 = l1.Next
//		}
//	}
//	return head.Next
//}
