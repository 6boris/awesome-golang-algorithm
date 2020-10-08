package Solution

func Solution(head *ListNode, G []int) int {
	GMap := make(map[int]bool)
	for _, val := range G {
		GMap[val] = true
	}
	ans, connected := 0, false
	for head != nil {
		if GMap[head.Val] {
			if !connected {
				ans++
			}
			connected = true
		} else {
			connected = false
		}
		head = head.Next
	}
	return ans
}
