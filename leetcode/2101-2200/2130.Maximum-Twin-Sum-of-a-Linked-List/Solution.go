package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solution(head *ListNode) int {
	ans := 0
	pairSum1(head, &head, &ans)
	return ans
}

func pairSum1(h1 *ListNode, h2 **ListNode, ans *int) {
	if h1.Next != nil {
		pairSum1(h1.Next, h2, ans)
	}

	h := *h2
	if r := h1.Val + h.Val; r > *ans {
		*ans = r
	}
	*h2 = h.Next
}
