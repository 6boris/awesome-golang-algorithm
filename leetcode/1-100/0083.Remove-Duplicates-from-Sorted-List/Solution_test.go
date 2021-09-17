package Solution

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs *ListNode
		expect *ListNode
	}{
		{
			"1 test 1",
			&ListNode{1, &ListNode{1, &ListNode{2, nil}}},
			&ListNode{1, &ListNode{2, nil}},
		},

		{
			"2 test 2",
			&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}},
			&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := deleteDuplicates(c.inputs)
			if !IsEqual(ret, c.expect) {
				PList(ret)
				PList(c.expect)
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)

			}
		})
	}
}

func PList(x *ListNode) {
	for x != nil {
		fmt.Print(x.Val, " ")
		x = x.Next
	}
	fmt.Println()
}

func IsEqual(x, y *ListNode) bool {
	for x != nil && y != nil {

		if x.Val != y.Val {
			return false
		}

		x = x.Next
		y = y.Next
	}

	if x != nil || y != nil {
		return false
	}

	return true
}
