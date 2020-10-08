package Solution

/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func Solution(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	p := new(ListNode)
	p.Next = head
	prev, current := p, p.Next
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			for current.Next != nil && current.Val == current.Next.Val {
				current = current.Next
			}
			current = current.Next
			prev.Next = current
		} else {
			prev = prev.Next
			current = current.Next
		}
	}
	return p.Next
}

