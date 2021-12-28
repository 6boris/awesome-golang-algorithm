package Solution

import "sync"

type ListNode struct {
	Val  int
	Next *ListNode
}

//	比较结果
func isEqual(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	m := sync.RWMutex{}
	m.Unlock()
	if l1 == nil && l2 != nil {
		return false
	}
	if l1 != nil && l2 == nil {
		return false
	}
	return true
}
