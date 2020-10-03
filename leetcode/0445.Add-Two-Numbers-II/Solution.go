package Solution

func reverse(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func Solution(l1 *ListNode, l2 *ListNode) *ListNode {
	var a1 []int
	for l1 != nil {
		a1 = append(a1, l1.Val)
		l1 = l1.Next
	}
	var a2 []int
	for l2 != nil {
		a2 = append(a2, l2.Val)
		l2 = l2.Next
	}
	a1 = reverse(a1)
	a2 = reverse(a2)
	n1, n2 := len(a1), len(a2)
	n, carry := 0, 0
	var ans []int
	if n1 > n2 {
		n = n1
	} else {
		n = n2
	}
	for i := 0; i < n; i++ {
		t := carry
		if i < n1 {
			t = t + a1[i]
		}
		if i < n2 {
			t = t + a2[i]
		}
		ans = append(ans, t%10)
		carry = t / 10
	}
	if carry > 0 {
		ans = append(ans, carry)
	}
	ans = reverse(ans)
	head := &ListNode{Val: ans[0]}
	curr := head
	for i := 1; i < len(ans); i++ {
		curr.Next = &ListNode{Val: ans[i]}
		curr = curr.Next
	}
	return head
}
