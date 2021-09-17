package Solution

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solution(head *ListNode) int {
	var bin strings.Builder
	for head != nil {
		bin.WriteString(strconv.Itoa(head.Val))
		head = head.Next
	}
	i, _ := strconv.ParseInt(bin.String(), 2, 64)
	return int(i)
}
