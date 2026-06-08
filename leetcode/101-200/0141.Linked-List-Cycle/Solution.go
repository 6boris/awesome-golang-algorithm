package Solution

import (
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 设置时间
func hasCycle1(head *ListNode) bool {
	start := time.Now()

	for time.Since(start) < 1000 {
		if head == nil {
			return false
		}
		// fmt.Println(head)

		head = head.Next
	}
	return true
}
