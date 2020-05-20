package Solution

func Solution(head *ListNode) []int {
	var nodes []int
	for head != nil {
		nodes = append(nodes, head.Val)
		head = head.Next
	}
	ans := make([]int, len(nodes))
	if len(nodes) == 0 {
		return ans
	}
	var stack []int
	for i := 0; i < len(nodes); i++ {
		for len(stack) > 0 && nodes[stack[len(stack)-1]] < nodes[i] {
			ans[stack[len(stack)-1]] = nodes[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
