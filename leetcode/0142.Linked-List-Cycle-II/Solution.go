package Solution

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	frist := head
	second := head
	isCycle := false

	for frist != nil && second != nil {
		frist = frist.Next
		if second.Next == nil {
			return nil
		}
		second = second.Next.Next
		if frist == second {
			isCycle = true
			break
		}
	}

	if !isCycle {
		return nil
	}
	frist = head

	for frist != second {
		frist = frist.Next
		second = second.Next
	}
	return frist

}
func Solution(x bool) bool {
	return x
}
