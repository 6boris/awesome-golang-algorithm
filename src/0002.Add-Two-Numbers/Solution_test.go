package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs [][]int
		expect bool
	}{
		{"1 test 1", [][]int{{2, 4, 3}, {5, 6, 4}}, true},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := addTwoNumbers()
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}

func MakeList(x []int) *ListNode {
	tmp := ListNode{Val: 0}
	head := tmp

	return head.Next
}
