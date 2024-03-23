package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solution(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	list2LastNode := list2
	for ; list2LastNode.Next != nil; list2LastNode = list2LastNode.Next {
	}

	var pre *ListNode
	cur := -1
	w := list1
	for w != nil {
		cur++
		if cur == a {
			if pre == nil {
				list1 = list2
			} else {
				pre.Next = list2
			}
		}
		if cur == b {
			list2LastNode.Next = w.Next
			break
		}
		pre = w
		w = w.Next
	}
	return list1
}
