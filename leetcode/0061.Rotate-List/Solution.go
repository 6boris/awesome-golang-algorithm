package Solution

func Solution(head *ListNode, k int) *ListNode {
	if nil == head || 0 == k {
		return head
	}
	head, length := cycleList(head)
	if k >= length {
		k = k % length
	}
	for i := 0; i < length-k-1; i++ {
		head = head.Next
	}
	p := head.Next
	head.Next = nil
	return p
}

func cycleList(l *ListNode) (*ListNode, int) {
	head, length := l, 1
	for l.Next != nil {
		l = l.Next
		length++
	}
	l.Next = head
	return head, length
}

func rotateRight(head *ListNode, k int) *ListNode {
	pointer := head
	count := 1
	if head == nil || head.Next == nil {
		return head
	}
	for pointer != nil {
		if pointer.Next == nil {
			pointer.Next = head
			break
		}
		pointer = pointer.Next
		count++
	}
	rem := count - k%count
	for rem > 0 {
		pointer = pointer.Next
		rem--
	}
	ans := pointer.Next
	pointer.Next = nil
	return ans
}
