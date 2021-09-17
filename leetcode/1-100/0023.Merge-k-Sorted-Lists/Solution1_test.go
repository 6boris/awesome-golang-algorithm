package Solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []*ListNode
		expect *ListNode
	}{
		{
			"TestCacse 1",
			[]*ListNode{
				UnmarshalListBySlice([]int{1, 4, 5}),
				UnmarshalListBySlice([]int{1, 3, 4}),
				UnmarshalListBySlice([]int{2, 6}),
			},
			UnmarshalListBySlice([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
	}
	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := mergeKLists3(c.inputs)
			if !isEqual(got, c.expect) {
				PrintList(got)
				PrintList(c.expect)
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}
