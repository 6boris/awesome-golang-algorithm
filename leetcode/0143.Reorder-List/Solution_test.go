package Solution

import "testing"

func TestSolution(t *testing.T) {
	// 测试用例
	cases := []struct {
		name string
		inputs []int
		expect *ListNode
	} {
		{
			"TestCase 1",
			[]int{1,2,3,4},
			UnmarshalListBySlice([]int{1,4,2,3}),
		},
		{
			"TestCase 2",
			[]int{1,2,3,4,5},
			UnmarshalListBySlice([]int{1,5,2,4,3}),
		},
	}

	// 开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			linkedList := UnmarshalListBySlice(c.inputs)
			Solution(linkedList)
			if !isEqual(linkedList, c.expect) {
				PrintList(linkedList)
				PrintList(c.expect)
				t.Fatalf("expected: %v, but got: %v, with inputs: %v", c.expect, linkedList, c.inputs)
			}
		})
	}
}
