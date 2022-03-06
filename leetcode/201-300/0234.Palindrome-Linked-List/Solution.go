package Solution

func Solution(head *ListNode) bool {
	walk := head
	r := true
	isPalindromeAux(head, &walk, &r)
	return r
}
func isPalindromeAux(head *ListNode, tail **ListNode, r *bool) {
	if head == nil {
		return
	}
	isPalindromeAux(head.Next, tail, r)
	res := head.Val == (*tail).Val

	*r = *r && res
	*tail = (*tail).Next
}
