package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func FindKthToTail(pListHead *ListNode, k int) *ListNode {
	if pListHead == nil || k <= 0 {
		return nil
	}
	n := 0
	cur := pListHead
	for cur != nil {
		cur = cur.Next
		n += 1
	}
	if n < k {
		return nil
	}
	n -= k

	for ; n > 0; n-- {
		pListHead = pListHead.Next

	}

	return pListHead
}

func FindKthToTail2(pListHead *ListNode, k int) *ListNode {
	if pListHead == nil || k <= 0 {
		return nil
	}
	slow, fast := pListHead, pListHead

	for ; k > 0; k-- {
		if fast != nil {
			fast = fast.Next
		} else {
			return nil
		}
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
