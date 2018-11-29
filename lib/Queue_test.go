package Solution

import (
	"fmt"
	"testing"
)

func TestQueue_Swap(t *testing.T) {
	que := Queue{}
	que.Push(&ListNode{Val: 1})
	fmt.Println(que.Pop().(*ListNode).Val)

}
