package Solution

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input1 *ListNode
		input2 *ListNode
		expect *ListNode
	}{
		{"1 test 1", &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}},
			&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}},
			&ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := addTwoNumbers(&ListNode{}, &ListNode{})

			for ret != nil {
				fmt.Println(ret.Val)
				ret = ret.Next
			}

			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input1)
			}
		})
	}
}
