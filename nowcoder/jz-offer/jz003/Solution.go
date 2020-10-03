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
	pre, next := &NodeList{}, &NodeList{}

	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}

	ans := []int{}

	for pre.Next != nil {
		ans = append(ans, pre.Val)
		pre = pre.Next
	}
	return ans

}
