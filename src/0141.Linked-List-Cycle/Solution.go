package Solution

import (
	"fmt"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//	设置时间
func hasCycle1(head *ListNode) bool {
	start := time.Now()

	for time.Since(start) < 1000 {
		if head == nil {
			return false
		}
		fmt.Println(head)

		head = head.Next
	}
	return true
}

//	HashMap
func hasCycle2(head *ListNode) bool {

	//listMap := make(map[string]int)

	//fmt.Println(listMap[fmt.Sprintf("%s",&head)])

	for head != nil {
		//fmt.Sprintf("%s",&head)

		//if listMap[fmt.Sprintf("%s", &head)] > 1 {
		//	//fmt.Println(listMap[fmt.Sprintf("%s", &head)])
		//	//fmt.Println(listMap)
		//	return true
		//}

		head = head.Next
	}
	return false
}

//	快慢指针
func hasCycle3(head *ListNode) bool {

	quick := head
	slow := head

	for quick.Next.Next != nil && slow.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next
		if slow == quick {
			return true
		}
	}
	return false
}
