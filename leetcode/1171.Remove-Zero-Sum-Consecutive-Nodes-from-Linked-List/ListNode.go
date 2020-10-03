package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeListNode(nodes []int) *ListNode {
	if len(nodes) == 0 {
		return &ListNode{}
	}
	list := &ListNode{}
	head := list
	list.Val = nodes[0]
	for i := 1; i < len(nodes); i++ {
		list.Next = &ListNode{}
		list = list.Next
		list.Val = nodes[i]
	}
	return head
}
