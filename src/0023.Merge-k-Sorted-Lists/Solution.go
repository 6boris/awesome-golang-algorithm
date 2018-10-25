package Solution

//	递归方法合并
func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	n := len(lists) / 2
	return merge(mergeKLists1(lists[:n]), mergeKLists1(lists[n:]))
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	node := head

	for {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				node.Next = l1
				node, l1 = node.Next, l1.Next
			} else {
				node.Next = l2
				node, l2 = node.Next, l2.Next
			}
		} else {
			if l1 == nil {
				node.Next = l2
			} else {
				node.Next = l1
			}
			break
		}
	}
	return head.Next
}
