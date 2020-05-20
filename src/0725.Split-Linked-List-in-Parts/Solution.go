package Solution

func Solution(root *ListNode, k int) []*ListNode {
	nodes := 0
	curr1 := root
	for curr1 != nil {
		nodes++
		curr1 = curr1.Next
	}
	ans := make([]*ListNode, k)
	quo, rem := nodes/k, nodes%k

	curr := root
	for i := 0; curr != nil && i < k; i++ {
		ans[i] = curr
		size := quo
		if rem > 0 {
			size++
			rem--
		}

		for j := 1; j < size; j++ {
			curr = curr.Next
		}
		end := curr
		curr = curr.Next
		end.Next = nil
	}
	return ans
}
