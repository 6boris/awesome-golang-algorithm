package Solution

import (
	"fmt"
	"testing"
)

func TestQueue_Swap(t *testing.T) {
	que := Queue{}
	que.Push(&ListNode{Val: 1})
	que.Push(&ListNode{Val: 2})
	fmt.Println(que.Len())
	fmt.Println(que.Less(1, 0))
	fmt.Println(que.Pop().(*ListNode).Val)
	que.Push(&ListNode{Val: 3})
	que.Push(&ListNode{Val: 4})

	que.Swap(0, 1)
}
