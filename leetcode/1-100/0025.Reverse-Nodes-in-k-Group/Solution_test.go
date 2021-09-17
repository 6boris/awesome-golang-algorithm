package Solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input1 *ListNode
		input2 int
		expect *ListNode
	}{
		{
			"TestCacse 1",
			UnmarshalListBySlice([]int{1, 2, 3, 4, 5}),
			2,
			UnmarshalListBySlice([]int{2, 1, 4, 3, 5}),
		},
		{
			"TestCacse 2",
			UnmarshalListBySlice([]int{1, 2, 3, 4, 5}),
			3,
			UnmarshalListBySlice([]int{3, 2, 1, 4, 5}),
		},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := reverseKGroup(c.input1, c.input2)
			if !isEqual(got, c.expect) {
				// t.Fatalf("expected: %v, but got: %v, with inputs: %v",
				// c.expect, got, c.input1)

				PrintList(got)
				PrintList(c.expect)

			}
		})
	}
}
