package Solution

type NodeList struct {
	Val  int
	Next *NodeList
}

//	递归写法
func printListFromTailToHead(head *NodeList) []int {
	ans := make([]int, 0)
	if head == nil {
		return ans
	}
	ans = printListFromTailToHead(head.Next)
	ans = append(ans, head.Val)

	return ans
}

//	反转琏表
func printListFromTailToHead2(head *NodeList) []int {
	pre, cur, next, ans := &NodeList{}, head, head.Next, []int{}

	for cur != nil {
		next = cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}

	for pre.Next != nil {
		ans = append(ans, pre.Val)
		pre = pre.Next
	}

	return ans
}
